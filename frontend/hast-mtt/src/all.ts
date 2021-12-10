import {one} from './one'
import {H, HastNode} from './types'

export function all(h: H, parent: HastNode): Array<any> {
  let nodes = parent.children || []
  let values = []
  let index = -1
  let length = nodes.length
  let child = nodes[index + 1]

  // Trim initial and final `<br>`s.
  // They’re not semantic per HTML
  while (child && child.type == 'element' && child.tagName == 'br') {
    index++
    child = nodes[index + 1]
  }

  child = nodes[length - 1]

  while (length - 1 > index && child && child.type == 'element' && child.tagName == 'br') {
    length--
    child = nodes[length - 1]
  }

  while (++index < length) {
    const result = one(h, nodes[index], parent)
    if (Array.isArray(result)) {
      values.push(...result.map(({position, ...n}) => n))
    } else if (result) {
      const {position, ...node} = result
      values.push(node)
    }
  }
  return values
}