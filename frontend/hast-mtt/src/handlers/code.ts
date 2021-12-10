import {code as buildCode, createId, paragraph, text} from '@mintter/mttast-builder'
import {convertElement} from 'hast-util-is-element'
import {toText} from 'hast-util-to-text'
import {trimTrailingLines} from 'trim-trailing-lines'
import {H, HastNode} from '../types'
import {wrapText} from '../util/wrap-text'

const prefix = 'language-'
let isPre = convertElement('pre')
let isCode = convertElement('code')

export function code(h: H, node: HastNode) {
  const children = node.children
  let index = -1
  let classList: Array<string | number | undefined>
  let lang: string | undefined

  if (isPre(node)) {
    while (++index < children.length) {
      const child = children[index]

      if (isCode(child) && child.properties?.className && Array.isArray(child.properties.className)) {
        classList = child.properties.className
        break
      }
    }
  }
  // @ts-ignore
  if (classList) {
    index = -1

    while (++index < classList.length) {
      if (String(classList[index]).slice(0, prefix.length) === prefix) {
        lang = String(classList[index]).slice(prefix.length)
        break
      }
    }
  }

  let props: any = {
    id: createId(),
  }

  if (lang) {
    props.lang = lang as any
  }

  return buildCode(props, [paragraph([text(trimTrailingLines(wrapText(h, toText(node))))])])
}