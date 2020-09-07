import {ImageKeyOption, ImagePluginOptionsValues} from '@udecode/slate-plugins'
import {ImageBlock} from './components/image'

export const ELEMENT_IMAGE = 'img'

export const DEFAULTS_IMAGE: Record<
  ImageKeyOption,
  Required<ImagePluginOptionsValues>
> = {
  img: {
    component: ImageBlock,
    type: ELEMENT_IMAGE,
    rootProps: {},
  },
}