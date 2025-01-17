import {darkTheme, lightTheme, styled} from '@app/stitches.config'

const Color = styled('div', {
  width: '$full',
  minWidth: 24,
  height: '$full',
  aspectRatio: '1',
  minHeight: 0,
  boxShadow: '0 0 1px black',
})

const ColorWrapper = styled('div', {
  display: 'grid',
  width: '$full',
  // gridTemplateColumns: 'repeat(auto-fill, minmax(24px, 1fr))',
  gridTemplateColumns: 'repeat(13, 1fr)',
  padding: '$3',
  gap: '$2',
  background: '$base-background-normal',
  boxSizing: 'border-box',
})

describe('Colors', () => {
  it('Default', () => {
    cy.mount(
      <>
        <Colors theme={lightTheme} />
        <Colors theme={darkTheme} />
      </>,
    )
  })
})

function Colors({theme}: {theme: typeof lightTheme | typeof darkTheme}) {
  return (
    <ColorWrapper>
      {Object.values(theme.colors).map(({value}) => (
        <Color key={value} css={{backgroundColor: value}} />
      ))}
    </ColorWrapper>
  )
}
