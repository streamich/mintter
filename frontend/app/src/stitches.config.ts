import {
  blue,
  blueDark,
  gray,
  grayDark,
  green,
  greenDark,
  orange,
  orangeDark,
  red,
  redDark,
  yellow,
  yellowDark,
} from '@radix-ui/colors'
import type * as Stitches from '@stitches/react'
import {createStitches, defaultThemeMap} from '@stitches/react'

const stitches = createStitches({
  prefix: 'mtt',
  theme: {
    borderStyles: {},
    borderWidths: {},
    colors: {
      'base-background-subtle': gray.gray1,
      'base-background-normal': gray.gray2,
      'base-component-bg-normal': gray.gray3,
      'base-component-bg-hover': gray.gray4,
      'base-component-bg-active': gray.gray5,
      'base-border-subtle': gray.gray6,
      'base-border-normal': gray.gray7,
      'base-border-hover': gray.gray8,
      'base-normal': gray.gray9,
      'base-active': gray.gray10,
      'base-text-low': gray.gray11,
      'base-text-high': gray.gray12,
      'base-text-opposite': gray.gray1,

      'primary-background-subtle': blue.blue1,
      'primary-background-normal': blue.blue2,
      'primary-component-bg-normal': blue.blue3,
      'primary-component-bg-hover': blue.blue4,
      'primary-component-bg-active': blue.blue5,
      'primary-border-subtle': blue.blue6,
      'primary-border-normal': blue.blue7,
      'primary-border-hover': blue.blue8,
      'primary-normal': blue.blue9,
      'primary-active': blue.blue10,
      'primary-text-low': blue.blue11,
      'primary-text-high': blue.blue12,
      'primary-text-opposite': blue.blue1,

      'secondary-background-subtle': orange.orange1,
      'secondary-background-normal': orange.orange2,
      'secondary-component-bg-normal': orange.orange3,
      'secondary-component-bg-hover': orange.orange4,
      'secondary-component-bg-active': orange.orange5,
      'secondary-border-subtle': orange.orange6,
      'secondary-border-normal': orange.orange7,
      'secondary-border-hover': orange.orange8,
      'secondary-normal': orange.orange9,
      'secondary-active': orange.orange10,
      'secondary-text-low': orange.orange11,
      'secondary-text-high': orange.orange12,
      'secondary-text-opposite': orange.orange1,

      'success-background-subtle': green.green1,
      'success-background-normal': green.green2,
      'success-component-bg-normal': green.green3,
      'success-component-bg-hover': green.green4,
      'success-component-bg-active': green.green5,
      'success-border-subtle': green.green6,
      'success-border-normal': green.green7,
      'success-border-hover': green.green8,
      'success-normal': green.green9,
      'success-active': green.green10,
      'success-text-low': green.green11,
      'success-text-high': green.green12,
      'success-text-opposite': green.green1,

      'warning-background-subtle': yellow.yellow1,
      'warning-background-normal': yellow.yellow2,
      'warning-component-bg-normal': yellow.yellow3,
      'warning-component-bg-hover': yellow.yellow4,
      'warning-component-bg-active': yellow.yellow5,
      'warning-border-subtle': yellow.yellow6,
      'warning-border-normal': yellow.yellow7,
      'warning-border-hover': yellow.yellow8,
      'warning-normal': yellow.yellow9,
      'warning-active': yellow.yellow10,
      'warning-text-low': yellow.yellow11,
      'warning-text-high': yellow.yellow12,
      'warning-text-opposite': yellow.yellow1,

      'danger-background-subtle': red.red1,
      'danger-background-normal': red.red2,
      'danger-component-bg-normal': red.red3,
      'danger-component-bg-hover': red.red4,
      'danger-component-bg-active': red.red5,
      'danger-border-subtle': red.red6,
      'danger-border-normal': red.red7,
      'danger-border-hover': red.red8,
      'danger-normal': red.red9,
      'danger-active': red.red10,
      'danger-text-low': red.red11,
      'danger-text-high': red.red12,
      'danger-text-opposite': red.red1,
    },
    fonts: {
      base: `-apple-system, BlinkMacSystemFont, 'Segoe UI', Helvetica, Arial,
      sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji'`,
      alt: `Iowan Old Style, Apple Garamond, Baskerville, Times New Roman, Droid
      Serif, Times, Source Serif Pro, serif, Apple Color Emoji, Segoe UI Emoji, Segoe
      UI Symbol`,
    },
    fontSizes: {
      1: '0.75rem',
      2: '0.85rem',
      3: '1rem',
      4: '1.33rem',
      5: '1.77rem',
      6: '2.36rem',
      7: '3.15rem',
      8: '4.2rem',
    },
    fontWeights: {
      regular: 400,
      medium: 500,
      bold: 700,
    },
    letterSpacings: {},
    lineHeights: {
      1: '1.1',
      2: '1.3',
      3: '1.5',
      4: '2',
    },
    radii: {
      1: '3px',
      2: '5px',
      3: '7px',
      round: '50%',
      pill: '9999px',
    },
    shadows: {
      focus: '0 0 0 2px $colors$primary-border-hover',
      menu: '0px 4px 8px $colors$base-component-bg-active',
      3: '0px 4px 8px $colors$base-component-bg-active',
      debug: '0 0 0 2px $colors$base-active',
    },
    sizes: {
      none: '0px',
      'one-quarter': '25%',
      'one-third': '33.333%',
      half: '50%',
      'two-thirds': '66.6666%',
      'three-quarters': '75%',
      full: '100%',
      'prose-width': '59ch',
      'library-width': 'var(--library-size)',
    },
    space: {
      0: 0,
      1: '2px',
      2: '4px',
      3: '8px',
      4: '12px',
      5: '16px',
      6: '20px',
      7: '24px',
      8: '32px',
      9: '64px',
    },
    zIndices: {
      1: '100',
      2: '200',
      3: '300',
      4: '400',
      max: '99999',
    },
  },
  media: {
    bp1: '(min-width: 768px)',
    bp2: '(min-width: 1024px)',
    bp3: '(min-width: 1400px)',
    dark: '(prefers-color-scheme: dark)',
    light: '(prefers-color-scheme: light)',
  },
  utils: {
    marginHorizontal: (val: Stitches.PropertyValue<'margin'>) => ({
      marginInline: val,
    }),
    marginVertical: (val: Stitches.PropertyValue<'margin'>) => ({
      marginBlock: val,
    }),
    paddingHorizontal: (val: Stitches.PropertyValue<'padding'>) => ({
      paddingInline: val,
    }),
    paddingVertical: (val: Stitches.PropertyValue<'padding'>) => ({
      paddingBlock: val,
    }),
  },
  themeMap: {
    ...defaultThemeMap,
    marginHorizontal: 'space' as const,
    marginVertical: 'space' as const,
    paddingHorizontal: 'space' as const,
    paddingVertical: 'space' as const,
  },
})

export const {styled, css, globalCss, keyframes, config, createTheme} = stitches
export type CSS = Stitches.CSS<typeof config>
export type ThemeColorValue = Stitches.PropertyValue<'backgroundColor'>

export const lightTheme = createTheme('light-theme', {
  colors: {
    'base-background-subtle': gray.gray1,
    'base-background-normal': gray.gray2,
    'base-component-bg-normal': gray.gray3,
    'base-component-bg-hover': gray.gray4,
    'base-component-bg-active': gray.gray5,
    'base-border-subtle': gray.gray6,
    'base-border-normal': gray.gray7,
    'base-border-hover': gray.gray8,
    'base-normal': gray.gray9,
    'base-active': gray.gray10,
    'base-text-low': gray.gray11,
    'base-text-high': gray.gray12,
    'base-text-opposite': gray.gray1,

    'primary-background-subtle': blue.blue1,
    'primary-background-normal': blue.blue2,
    'primary-component-bg-normal': blue.blue3,
    'primary-component-bg-hover': blue.blue4,
    'primary-component-bg-active': blue.blue5,
    'primary-border-subtle': blue.blue6,
    'primary-border-normal': blue.blue7,
    'primary-border-hover': blue.blue8,
    'primary-normal': blue.blue9,
    'primary-active': blue.blue10,
    'primary-text-low': blue.blue11,
    'primary-text-high': blue.blue12,
    'primary-text-opposite': blue.blue1,

    'secondary-background-subtle': orange.orange1,
    'secondary-background-normal': orange.orange2,
    'secondary-component-bg-normal': orange.orange3,
    'secondary-component-bg-hover': orange.orange4,
    'secondary-component-bg-active': orange.orange5,
    'secondary-border-subtle': orange.orange6,
    'secondary-border-normal': orange.orange7,
    'secondary-border-hover': orange.orange8,
    'secondary-normal': orange.orange9,
    'secondary-active': orange.orange10,
    'secondary-text-low': orange.orange11,
    'secondary-text-high': orange.orange12,
    'secondary-text-opposite': orange.orange1,

    'success-background-subtle': green.green1,
    'success-background-normal': green.green2,
    'success-component-bg-normal': green.green3,
    'success-component-bg-hover': green.green4,
    'success-component-bg-active': green.green5,
    'success-border-subtle': green.green6,
    'success-border-normal': green.green7,
    'success-border-hover': green.green8,
    'success-normal': green.green9,
    'success-active': green.green10,
    'success-text-low': green.green11,
    'success-text-high': green.green12,
    'success-text-opposite': green.green1,

    'warning-background-subtle': yellow.yellow1,
    'warning-background-normal': yellow.yellow2,
    'warning-component-bg-normal': yellow.yellow3,
    'warning-component-bg-hover': yellow.yellow4,
    'warning-component-bg-active': yellow.yellow5,
    'warning-border-subtle': yellow.yellow6,
    'warning-border-normal': yellow.yellow7,
    'warning-border-hover': yellow.yellow8,
    'warning-normal': yellow.yellow9,
    'warning-active': yellow.yellow10,
    'warning-text-low': yellow.yellow11,
    'warning-text-high': yellow.yellow12,
    'warning-text-opposite': yellow.yellow1,

    'danger-background-subtle': red.red1,
    'danger-background-normal': red.red2,
    'danger-component-bg-normal': red.red3,
    'danger-component-bg-hover': red.red4,
    'danger-component-bg-active': red.red5,
    'danger-border-subtle': red.red6,
    'danger-border-normal': red.red7,
    'danger-border-hover': red.red8,
    'danger-normal': red.red9,
    'danger-active': red.red10,
    'danger-text-low': red.red11,
    'danger-text-high': red.red12,
    'danger-text-opposite': red.red1,
  },
})

export const darkTheme = createTheme('dark-theme', {
  colors: {
    'base-background-subtle': grayDark.gray1,
    'base-background-normal': grayDark.gray2,
    'base-component-bg-normal': grayDark.gray3,
    'base-component-bg-hover': grayDark.gray4,
    'base-component-bg-active': grayDark.gray5,
    'base-border-subtle': grayDark.gray6,
    'base-border-normal': grayDark.gray7,
    'base-border-hover': grayDark.gray8,
    'base-normal': grayDark.gray9,
    'base-active': grayDark.gray10,
    'base-text-low': grayDark.gray11,
    'base-text-high': grayDark.gray12,
    'base-text-opposite': grayDark.gray1,

    'primary-background-subtle': blueDark.blue1,
    'primary-background-normal': blueDark.blue2,
    'primary-component-bg-normal': blueDark.blue3,
    'primary-component-bg-hover': blueDark.blue4,
    'primary-component-bg-active': blueDark.blue5,
    'primary-border-subtle': blueDark.blue6,
    'primary-border-normal': blueDark.blue7,
    'primary-border-hover': blueDark.blue8,
    'primary-normal': blueDark.blue9,
    'primary-active': blueDark.blue10,
    'primary-text-low': blueDark.blue11,
    'primary-text-high': blueDark.blue12,
    'primary-text-opposite': blueDark.blue1,

    'secondary-background-subtle': orangeDark.orange1,
    'secondary-background-normal': orangeDark.orange2,
    'secondary-component-bg-normal': orangeDark.orange3,
    'secondary-component-bg-hover': orangeDark.orange4,
    'secondary-component-bg-active': orangeDark.orange5,
    'secondary-border-subtle': orangeDark.orange6,
    'secondary-border-normal': orangeDark.orange7,
    'secondary-border-hover': orangeDark.orange8,
    'secondary-normal': orangeDark.orange9,
    'secondary-active': orangeDark.orange10,
    'secondary-text-low': orangeDark.orange11,
    'secondary-text-high': orangeDark.orange12,
    'secondary-text-opposite': orangeDark.orange1,

    'success-background-subtle': greenDark.green1,
    'success-background-normal': greenDark.green2,
    'success-component-bg-normal': greenDark.green3,
    'success-component-bg-hover': greenDark.green4,
    'success-component-bg-active': greenDark.green5,
    'success-border-subtle': greenDark.green6,
    'success-border-normal': greenDark.green7,
    'success-border-hover': greenDark.green8,
    'success-normal': greenDark.green9,
    'success-active': greenDark.green10,
    'success-text-low': greenDark.green11,
    'success-text-high': greenDark.green12,
    'success-text-opposite': greenDark.green1,

    'warning-background-subtle': yellowDark.yellow1,
    'warning-background-normal': yellowDark.yellow2,
    'warning-component-bg-normal': yellowDark.yellow3,
    'warning-component-bg-hover': yellowDark.yellow4,
    'warning-component-bg-active': yellowDark.yellow5,
    'warning-border-subtle': yellowDark.yellow6,
    'warning-border-normal': yellowDark.yellow7,
    'warning-border-hover': yellowDark.yellow8,
    'warning-normal': yellowDark.yellow9,
    'warning-active': yellowDark.yellow10,
    'warning-text-low': yellowDark.yellow11,
    'warning-text-high': yellowDark.yellow12,
    'warning-text-opposite': yellowDark.yellow1,

    'danger-background-subtle': redDark.red1,
    'danger-background-normal': redDark.red2,
    'danger-component-bg-normal': redDark.red3,
    'danger-component-bg-hover': redDark.red4,
    'danger-component-bg-active': redDark.red5,
    'danger-border-subtle': redDark.red6,
    'danger-border-normal': redDark.red7,
    'danger-border-hover': redDark.red8,
    'danger-normal': redDark.red9,
    'danger-active': redDark.red10,
    'danger-text-low': redDark.red11,
    'danger-text-high': redDark.red12,
    'danger-text-opposite': redDark.red1,
  },
})

/**
 * some of the reset styles down here are based on this links:
 * - https://www.joshwcomeau.com/css/custom-css-reset/
 *
 */

export const globalStyles = globalCss({
  ':root': {
    '--library-size': '232px',
    '--tools-x': '-999px',
    '--tools-y': '-999px',
    '--block-h': '0',
    '--topbar-h': '40px',
    '--footer-h': '24px',
  },
  '*': {
    boxSizing: 'border-box',
    // margin: 0,
  },
  html: {
    blockSize: '100%',
  },
  body: {
    backgroundColor: '$base-background-subtle',
    color: '$base-text-high',
    lineHeight: '$3',
    fontFamily: '$base',
    fontSize: '100%',
    width: '100%',
    height: '100%',
    overflow: 'hidden',
    textRendering: 'optimizeSpeed',
    '-webkit-font-smoothing': 'antialiased',
    '-moz-osx-font-smoothing': 'grayscale',
  },
  'a:not([class])': {
    color: 'currentColor',
  },
  'main:focus': {
    outline: 'none',
  },
  '[href]:hover': {
    cursor: 'pointer',
    textDecoration: 'underline',
  },
  'p, h1, h2, h3, h4, h5, h6': {
    overflowWrap: 'break-word',
  },
  'h1, h2, h3, h4': {
    fontFamily: '$base',
    lineHeight: '$2',
  },
  'h1, h2': {
    fontWeight: '$bold',
  },
  h1: {
    fontSize: '$5',
    '@bp1': {
      fontSize: '$6',
    },
  },
  h2: {
    fontSize: '$4',
    '@bp1': {
      fontSize: '$5',
    },
  },
  'h3,h4,h5,h6': {fontSize: '$4', marginBlock: '1rem'},
  'input, button, textarea, select': {
    font: 'inherit',
  },
  img: {
    maxWidth: 'min(55rem, 100%)', // review
  },
  figcaption: {
    fontSize: '$1',
    fontStyle: 'italic',
    marginTop: '$5',
    '@bp1': {
      fontSize: '$2',
    },
  },
  p: {
    fontFamily: '$alt',
    fontSize: '$3',
    '@bp1': {
      // fontSize: '$4',
    },
  },
  blockquote: {
    marginInlineStart: 0,
    paddingInlineStart: '$5',
    borderInlineStart: '3px solid $colors$primary-active',
    fontStyle: 'italic',
    '& p': {
      fontSize: '$4',
      maxWidth: '40ch',
      color: '$base-text-low',
    },
  },
  pre: {
    // maxWidth: '80ch',
    backgroundColor: '$base-background-normal',
    paddingInlineStart: '2rem',
    paddingBlock: '$5',
    marginInlineEnd: '$4',
    fontSize: '$4',
    whiteSpace: 'pre',
  },
  '::selection': {
    backgroundColor: '$primary-active',
    color: 'white',
  },
  '@dark': {
    // notice the `media` definition on the stitches.config.ts file
    ':root:not(.light)': {
      ...Object.keys(darkTheme.colors).reduce((varSet, currentColorKey) => {
        const currentColor = darkTheme.colors[currentColorKey]
        const currentColorValue =
          currentColor.value.substring(0, 1) === '$'
            ? `$colors${currentColor.value}`
            : currentColor.value

        return {
          [currentColor.variable]: currentColorValue,
          ...varSet,
        }
      }, {}),
    },
  },
})

export const flow = css({
  '& li > * + li *': {
    marginTop: 'var(--flow-space, 1em)',
  },
})
