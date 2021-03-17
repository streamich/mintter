import React from 'react'
import {
  LinkKeyOption,
  LinkPluginOptionsValues,
  isUrl,
} from '@udecode/slate-plugins'
import {css} from 'emotion'
import {ELEMENT_LINK} from './defaults'

export const DEFAULTS_LINK: Record<LinkKeyOption, LinkPluginOptionsValues> = {
  link: {
    component: ({element, attributes, children, as: Component = 'a'}) =>
      element.type === ELEMENT_LINK ? (
        <Component
          {...attributes}
          className={`text-primary cursor-pointer hover:text-primary-hover transition duration-200 hover:cursor-pointer ${css`
            p {
              display: inline;

              &:hover {
                cursor: pointer;
              }
            }
          `}`}
          onClick={() => window.open(element.url as string, '_blank')}
          href={element.url as string}
        >
          {children}
        </Component>
      ) : null,
    type: ELEMENT_LINK,
    rootProps: {
      as: 'a',
    },
    isUrl,
  },
}