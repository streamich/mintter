import {
  MarkOnKeyDownOptions,
  StrikethroughKeyOption,
  StrikethroughPluginOptionsValues,
} from '@udecode/slate-plugins'

export const MARK_STRIKETHROUGH = 'strikethrough'

export const STRIKETHROUGH_OPTIONS: Record<
  StrikethroughKeyOption,
  StrikethroughPluginOptionsValues & MarkOnKeyDownOptions
> = {
  strikethrough: {
    type: MARK_STRIKETHROUGH,
    hotkey: 'mod+shift+s',
    rootProps: {
      className: `line-through`,
      as: 's',
    },
  },
}