import {Path, Transforms} from 'slate'
import type {EditorPlugin} from '../types'
import {Editor} from 'slate'
import {isGroupContent, isHeading, isStatement, isText} from '@mintter/mttast'
import type {Statement as StatementType} from '@mintter/mttast'
import {isLastChild, getLastChild} from '../utils'
import type {MTTEditor} from '../utils'
import {styled} from '@mintter/ui/stitches.config'
import {group} from '@mintter/mttast-builder'
import {Icon} from '@mintter/ui/icon'
import {Text} from '@mintter/ui/text'
import {Marker} from '../marker'
import {isFirstChild} from '@udecode/slate-plugins'
import type {NodeEntry} from 'slate'
import * as ContextMenu from '@radix-ui/react-context-menu'
import {useParams} from 'react-router'
import toast from 'react-hot-toast'
import {send} from 'xstate'
import {useSidepanel} from '../../components/sidepanel'

export const ELEMENT_STATEMENT = 'statement'

export const Tools = styled('div', {
  height: '$space$8',
  overflow: 'hidden',
  alignSelf: 'center',
  display: 'flex',
  alignItems: 'center',
})
const StatementStyled = styled('li', {
  marginTop: '$3',
  padding: 0,
  listStyle: 'none',
  display: 'grid',
  wordBreak: 'break-word',
  gridTemplateColumns: 'minmax($space$8, auto) 1fr',
  gridTemplateRows: 'min-content auto',
  gap: '0 $2',
  gridTemplateAreas: `"controls content"
  ". children"`,
  [`& > ${Tools}`]: {
    gridArea: 'controls',
  },
  "& > [data-element-type='paragraph']": {
    gridArea: 'content',
  },
  '& > ul, & > ol': {
    gridArea: 'children',
  },
})

export const Dragger = styled('div', {
  // backgroundColor: 'red',
  width: '$space$8',
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'center',
  height: 32,
  borderRadius: '$2',
  opacity: 0,
  transition: 'all ease-in-out 0.1s',
  '&:hover': {
    opacity: 1,
    cursor: 'grab',
  },
})

const ContextMenuContentStyled = styled(ContextMenu.Content, {
  minWidth: 130,
  backgroundColor: 'white',
  borderRadius: 6,
  padding: 5,
  boxShadow: '0px 5px 15px -5px hsla(206,22%,7%,.15)',
})

const StyledItem = styled(ContextMenu.Item, {
  display: 'flex',
  flexDirection: 'row',
  alignItems: 'center',
  justifyContent: 'start',
  gap: '$4',
  paddingVertical: '$2',
  paddingHorizontal: '$4',
  cursor: 'pointer',
  '&:focus': {
    outline: 'none',
    backgroundColor: '$primary-muted',
    cursor: 'pointer',
  },
})

export const createStatementPlugin = (): EditorPlugin => ({
  name: ELEMENT_STATEMENT,
  renderElement({attributes, children, element}) {
    // TODO: create a hook to get all the embeds from its content to rendered as annotations. this should be used to all FlowContent nodes
    // const editor = useSlateStatic()
    // const currentPath = ReactEditor.findPath(editor, element)
    // const [parent, parentPath] = Editor.parent(editor, currentPath)
    // useEffect(() => {
    //   console.log('parent new childs!', element, parent)
    // }, [parent.children.length])
    if (isStatement(element)) {
      const {docId} = useParams<{docId: string}>()
      const {send} = useSidepanel()
      async function onCopy() {
        await copyTextToClipboard(`mtt://${docId}/${(element as StatementType).id}`)
        toast.success('Statement Reference copied successfully', {position: 'top-center'})
      }
      function onStartDraft() {}
      return (
        <StatementStyled {...attributes}>
          <Tools contentEditable={false}>
            <Dragger data-dragger>
              <Icon name="Grid6" size="2" color="muted" />
            </Dragger>
            <Marker element={element} />
          </Tools>
          <ContextMenu.Root>
            <ContextMenu.Trigger>{children}</ContextMenu.Trigger>
            <ContextMenuContentStyled alignOffset={-5}>
              <StyledItem onSelect={onCopy}>
                <Icon name="Copy" size="1" />
                <Text size="2">Copy Statement Reference</Text>
              </StyledItem>
              <StyledItem onSelect={() => send({type: 'SIDEPANEL_ADD_ITEM', payload: `${docId}/${element.id}`})}>
                <Icon size="1" name="ArrowBottomRight" />
                <Text size="2">Open in Sidepanel</Text>
              </StyledItem>
              <StyledItem onSelect={onStartDraft}>
                <Icon size="1" name="AddCircle" />
                <Text size="2">Start a Draft</Text>
              </StyledItem>
            </ContextMenuContentStyled>
          </ContextMenu.Root>
        </StatementStyled>
      )
    }
  },
  configureEditor(editor) {
    const {normalizeNode} = editor

    editor.normalizeNode = (entry) => {
      const [node, path] = entry
      if (isStatement(node)) {
        if (removeEmptyStatement(editor, entry)) return
        // check if there's a group below, if so, move inside that group
        const parent = Editor.parent(editor, path)

        if (!isLastChild(parent, path)) {
          const lastChild = getLastChild(parent)
          if (isGroupContent(lastChild[0])) {
            // the last child of the statement is a group. we should move the new as the first child
            console.log('move node!')
            Transforms.moveNodes(editor, {at: path, to: lastChild[1].concat(0)})
            return
          }
        }

        const [parentNode, parentPath] = parent
        if (isStatement(parentNode)) {
          // if parent is a statement and is the last child (because the previous if is false) then we can move the new statement to the next position of it's parent
          Transforms.moveNodes(editor, {
            at: path,
            to: isFirstChild(path) ? parentPath : Path.next(parentPath),
          })
          return
        }
        if (isHeading(parentNode)) {
          // this statement should be part of a group inside the heading, we need to wrap it!
          Transforms.wrapNodes(editor, group([]), {at: path})
          return
        }

        if (isGroupContent(node.children[0])) {
          Transforms.unwrapNodes(editor, {at: path})
          return
        }
      }
      normalizeNode(entry)
    }

    return editor
  },
})

export function removeEmptyStatement(editor: MTTEditor, entry: NodeEntry<StatementType>): boolean | undefined {
  const [node, path] = entry
  if (isStatement(node)) {
    if (node.children.length == 1) {
      const children = Editor.node(editor, path.concat(0))
      if (!('type' in children[0])) {
        Transforms.removeNodes(editor, {
          at: path,
        })
        return true
      }
    }
  }
}

function copyTextToClipboard(text: string) {
  return new Promise((resolve, reject) => {
    if (!navigator.clipboard) {
      return fallbackCopyTextToClipboard(text)
    }
    return navigator.clipboard.writeText(text).then(
      () => {
        resolve(text)
      },
      (err) => {
        console.error('Async: Could not copy text: ', err)
        reject(err)
      },
    )
  })
}

function fallbackCopyTextToClipboard(text: string) {
  return new Promise((resolve, reject) => {
    const textArea = document.createElement('textarea')
    textArea.value = text

    // Avoid scrolling to bottom
    textArea.style.top = '0'
    textArea.style.left = '0'
    textArea.style.position = 'fixed'
    textArea.style.opacity = '0'

    document.body.appendChild(textArea)
    textArea.focus()
    textArea.select()

    try {
      document.execCommand('copy')
    } catch (err) {
      console.error('Fallback: Oops, unable to copy', err)
      reject(err)
    }

    document.body.removeChild(textArea)
    resolve(true)
  })
}