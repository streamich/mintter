import React from 'react'

import {
  ParagraphKeyOption,
  ParagraphPluginOptionsValues,
} from '@udecode/slate-plugins'

export const ELEMENT_PARAGRAPH = 'p'

export const paragraphOption = {}

export const PARAGRAPH_OPTIONS: Record<
  ParagraphKeyOption,
  Required<ParagraphPluginOptionsValues>
> = {
  p: {
    component: ({
      as: Component = 'p',
      className = '',
      element,
      attributes,
      ...rest
    }) => {
      return element.type === ELEMENT_PARAGRAPH ? (
        <Component {...attributes} className={className} {...rest} />
      ) : null
    },
    type: ELEMENT_PARAGRAPH,
    rootProps: {
      className: 'text-body text-xl leading-loose',
      as: 'p',
    },
  },
}