import {css} from '@app/stitches.config'
import type {EditorPlugin} from '../types'

export const MARK_CODE = 'code'

const codeStyles = css({
  padding: '3px 5px',
  borderRadius: '$2',
  backgroundColor: '$base-component-bg-normal',
  color: '$base-text-hight',
  fontSize: '0.9em',
})

export const createInlineCodePlugin = (): EditorPlugin => ({
  name: MARK_CODE,
  renderLeaf:
    () =>
    ({attributes, children, leaf}) => {
      if (leaf[MARK_CODE] && leaf.value) {
        return (
          <code className={codeStyles()} {...attributes}>
            {children}
          </code>
        )
      }
    },
})
