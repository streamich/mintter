import {useBlockProps} from '@app/editor/editor-node-props'
import {MintterEditor} from '@app/editor/mintter-changes/plugin'
import {
  createId,
  FlowContent,
  isEmbed,
  isFlowContent,
  isGroupContent,
  isParagraph,
  isPhrasingContent,
  isStatement,
  isStaticParagraph,
  paragraph,
  statement,
  Statement as StatementType,
  text,
} from '@app/mttast'
import {Editor, Node, NodeEntry, Path, Transforms} from 'slate'
import {RenderElementProps} from 'slate-react'
import {EditorMode} from '../plugin-utils'
import type {EditorPlugin} from '../types'
import {isFirstChild, useBlockFlash} from '../utils'

export const ELEMENT_STATEMENT = 'statement'

export const createStatementPlugin = (): EditorPlugin => ({
  name: ELEMENT_STATEMENT,
  renderElement:
    (editor) =>
    ({element, children, attributes}) => {
      if (isStatement(element)) {
        return (
          <Statement
            mode={editor.mode}
            element={element}
            attributes={attributes}
          >
            {children}
          </Statement>
        )
      }
    },
  configureEditor(editor) {
    const {normalizeNode, insertBreak, deleteBackward} = editor

    editor.normalizeNode = (entry) => {
      const [node, path] = entry

      if (isStatement(node)) {
        // if the second child is also a paragraph and the third is a group,
        // move the second paragraph into the group
        if (
          addParagraphToNestedGroup(editor, entry as NodeEntry<StatementType>)
        ) {
          return
        }

        for (const [child, childPath] of Node.children(editor, path, {
          reverse: true,
        })) {
          if (isFirstChild(childPath)) {
            if (isFlowContent(child)) {
              Transforms.unwrapNodes(editor, {at: childPath})
              return
            }
            if (isStaticParagraph(child)) {
              Editor.withoutNormalizing(editor, () => {
                let {children} = child
                Transforms.insertNodes(editor, paragraph(children), {
                  at: childPath,
                })
                Transforms.removeNodes(editor, {at: Path.next(childPath)})
              })
            }

            if (isPhrasingContent(child)) {
              Transforms.wrapNodes(editor, paragraph([]), {at: childPath})
              return
            }
          }

          // move any flow content nodes that are children *outside* of this statement as siblings
          if (isFlowContent(child)) {
            Transforms.moveNodes(editor, {at: childPath, to: Path.next(path)})
            return
          }

          // if this is the second child & and it's a paragraph
          if (childPath[childPath.length - 1] == 1) {
            if (isParagraph(child)) {
              let index = childPath[childPath.length - 1]
              let nextChild = node.children[index + 1]
              // if the next child is a group
              // move this child into the group
              if (isGroupContent(nextChild)) {
                Transforms.moveNodes(editor, {
                  at: childPath,
                  to: Path.next(childPath).concat(0),
                })
                return
              } else {
                // else we move it outside the statement
                Transforms.moveNodes(editor, {
                  at: childPath,
                  to: Path.next(path),
                })
                return
              }
            }
          }

          // if this is the third or higher child
          // we move it outside the statement
          if (childPath[childPath.length - 1] > 1) {
            Transforms.moveNodes(editor, {at: childPath, to: Path.next(path)})
            return
          }

          // if the second child is a group, but the previous is a statement
          // move the group into the statement
          if (isGroupContent(child)) {
            let prev = Editor.previous(editor, {
              at: childPath,
              match: isFlowContent,
            })
            if (prev) {
              let [, pPath] = prev
              Transforms.moveNodes(editor, {at: childPath, to: pPath.concat(1)})
              return
            }
          }
        }
      }

      normalizeNode(entry)
    }

    editor.insertBreak = function blockInsertBreak() {
      let {selection} = editor

      // we need to run this code when the selection starts at the beginning of any node (usually statement or paragraph).
      // if it's in the beginning (no matter if its collapsed or not) we need to insert above instead of keeping the blockId in place

      if (selection?.anchor.offset == 0) {
        let currentEntry = Editor.above(editor, {
          match: isFlowContent,
        })

        if (currentEntry) {
          let [, path] = currentEntry
          let isEnd = Editor.isEnd(editor, selection.focus, path)
          if (isEnd) {
            insertBreak()
          } else {
            let newBlock = statement({id: createId()}, [paragraph([text('')])])
            Transforms.insertNodes(editor, newBlock, {at: path})
            MintterEditor.addChange(editor, ['moveBlock', newBlock.id])
            MintterEditor.addChange(editor, ['replaceBlock', newBlock.id])
            return
          }
        }
      } else {
        insertBreak()
      }
    }

    editor.deleteBackward = function blockDeleteBackwards(unit) {
      let {selection} = editor

      if (selection?.anchor.offset == 0) {
        let [node, path] =
          Editor.above(editor, {
            match: isFlowContent,
          }) || []

        if (node && path) {
          if (!isFirstChild(path)) {
            let prevBlockPath = Path.previous(path)
            let prevBlockNode = Node.get(editor, prevBlockPath)

            if (
              !Node.string(prevBlockNode) &&
              isFlowContent(prevBlockNode) &&
              !hasEmbedOnly([prevBlockNode, prevBlockPath])
            ) {
              Transforms.removeNodes(editor, {at: prevBlockPath})
              return
            }
          } else {
            /**
             * TODO:
             * here we need to make sure we are removing the current block so we are not leaving with orphan statements after we merge the content of both paragraphs (parent block paragraph and current paragraph)
             */
          }
        }
      }

      deleteBackward(unit)
    }

    return editor
  },
  onKeyDown: (editor) => (event) => {
    if (editor.selection && event.key == 'Enter') {
      if (event.shiftKey) {
        event.preventDefault()
        Transforms.insertText(editor, '\n')
        return
      }
    }
  },
})

function addParagraphToNestedGroup(
  editor: Editor,
  entry: NodeEntry<StatementType>,
): boolean | undefined {
  let [node, path] = entry
  if (
    node.children.length > 2 &&
    isParagraph(node.children[1]) &&
    isGroupContent(node.children[2])
  ) {
    Transforms.moveNodes(editor, {at: path.concat(1), to: path.concat(2, 0)})
    return true
  }
}

function Statement({
  attributes,
  children,
  element,
  mode,
}: RenderElementProps & {mode: EditorMode}) {
  let {blockProps} = useBlockProps(element)

  let inRoute = useBlockFlash(attributes.ref, element.id)

  if (mode == EditorMode.Embed) {
    return (
      <span {...attributes} {...blockProps}>
        {children}
      </span>
    )
  }

  return (
    <li
      {...attributes}
      {...blockProps}
      className={inRoute ? 'flash' : undefined}
    >
      {children}
    </li>
  )
}

export function removeEmptyStatement(
  editor: Editor,
  entry: NodeEntry<StatementType>,
): boolean | undefined {
  const [node, path] = entry

  if (Node.string(node).length == 0) {
    Transforms.removeNodes(editor, {
      at: path,
    })
    return true
  }
}

function hasEmbedOnly(entry: NodeEntry<FlowContent>) {
  let [node] = entry
  let hasContent = !!Node.string(node)
  let result = false

  if (!hasContent) {
    for (let childEntry of Node.descendants(node)) {
      let [child] = childEntry

      if (isEmbed(child)) {
        result = true
      }
    }
  }
  return result
}
