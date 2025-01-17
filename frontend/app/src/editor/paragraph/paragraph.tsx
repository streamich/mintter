import {usePhrasingProps} from '@app/editor/editor-node-props'
import {useBlockObserve, useMouse} from '@app/mouse-context'

import {
  isBlockquote,
  isCode,
  isParagraph,
  isPhrasingContent,
  Paragraph as ParagraphType,
} from '@app/mttast'
import {Box} from '@components/box'
import {Node, Path, Transforms} from 'slate'
import {RenderElementProps, useSlateStatic} from 'slate-react'
import {EditorMode} from '../plugin-utils'
import type {EditorPlugin} from '../types'

export const ELEMENT_PARAGRAPH = 'paragraph'

export const createParagraphPlugin = (): EditorPlugin => ({
  name: ELEMENT_PARAGRAPH,
  renderElement:
    (editor) =>
    ({element, children, attributes}) => {
      if (isParagraph(element)) {
        return (
          <Paragraph
            mode={editor.mode}
            element={element}
            attributes={attributes}
          >
            {children}
          </Paragraph>
        )
      }
    },
  configureEditor: (editor) => {
    const {normalizeNode} = editor

    editor.normalizeNode = (entry) => {
      const [node, path] = entry

      if (isParagraph(node)) {
        for (const [child, childPath] of Node.children(editor, path)) {
          if (!isPhrasingContent(child)) {
            Transforms.moveNodes(editor, {at: childPath, to: Path.next(path)})
            return
          }
        }
      }

      normalizeNode(entry)
    }

    return editor
  },
})

function Paragraph({
  children,
  element,
  attributes,
  mode,
}: RenderElementProps & {mode: EditorMode; element: ParagraphType}) {
  let editor = useSlateStatic()
  let {elementProps, parentNode} = usePhrasingProps(editor, element)

  useBlockObserve(mode, attributes.ref)
  let mouseService = useMouse()

  let mouseProps =
    mode != EditorMode.Discussion
      ? {
          onMouseEnter: () => {
            mouseService.send({
              type: 'HIGHLIGHT.ENTER',
              ref: elementProps['data-highlight'] as string,
            })
          },
          onMouseLeave: () => {
            mouseService.send('HIGHLIGHT.LEAVE')
          },
        }
      : {}

  if (mode == EditorMode.Embed) {
    return (
      <Box as="span" {...attributes} {...elementProps}>
        {children}
      </Box>
    )
  }

  if (isCode(parentNode)) {
    return (
      <Box as="pre" {...attributes} {...elementProps} {...mouseProps}>
        <code>{children}</code>
      </Box>
    )
  }

  if (isBlockquote(parentNode)) {
    return (
      <Box as="blockquote" {...attributes} {...elementProps} {...mouseProps}>
        <p>{children}</p>
      </Box>
    )
  }

  return (
    <p {...attributes} {...elementProps} {...mouseProps}>
      {children}
    </p>
  )
}
