import {blockquote as buildBlockquote, createId, paragraph, Text} from '../..'
import {all} from '../all'
import {H, HastNode} from '../types'
import {unwrap} from '../util/unwrap'

export function blockquote(h: H, node: HastNode) {
  const children = all(h, node)
  const content = unwrap<Text>(children, 'text')

  return buildBlockquote({id: createId()}, [paragraph(content)])
}
