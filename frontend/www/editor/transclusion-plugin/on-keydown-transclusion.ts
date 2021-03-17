import {isSelectionInTransclusion} from '../mintter-plugin/is-selection-in-block-item'
import {Editor, Path, Transforms} from 'slate'
import {getBlockAbove, getNextSiblingNodes} from '@udecode/slate-plugins'
import {id} from '../id'

export const onKeyDownTransclusion = options => (
  e: KeyboardEvent,
  editor: Editor,
) => {
  if (e.key === 'Enter') {
    const res = isSelectionInTransclusion(editor, options)
    if (!res) return
    createEmptyBlock(editor, options, res.blockPath)
    return
  }

  if (e.key === 'ArrowDown') {
    const res = isSelectionInTransclusion(editor, options)
    if (!res) return
    const blockAbove = getBlockAbove(editor, {at: res.blockPath})
    const [nextBlock] = getNextSiblingNodes(blockAbove, res.blockPath)

    if (!nextBlock) {
      createEmptyBlock(editor, options, res.blockPath)
    }
  }
}

function createEmptyBlock(editor: Editor, options, after) {
  const nextPath = Path.next(after)

  Transforms.insertNodes(
    editor,
    {
      type: options.block.type,
      id: id(),
      children: [{type: options.p.type, children: [{text: ''}]}],
    },
    {
      at: nextPath,
    },
  )
  Transforms.select(editor, Editor.start(editor, nextPath))
}