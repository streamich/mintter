import {MainPage} from '@app/pages/main-page'
import {mountProviders} from '@app/test/utils'

// TODO: FIXME
describe.skip('DraftList', () => {
  // TODO: maybe there are two mainServices started here, I'm getting DraftList and PubList queryClient errors (hitting the )
  it('Should show an empty list', () => {
    const {client, render, mainService} = mountProviders({
      initialRoute: '/drafts',
      draftList: [],
      publicationList: [],
    })

    console.log('TEST MAIN SERVICE', mainService.getSnapshot())

    render(<MainPage mainService={mainService} />).get(
      '[data-testid="topbar-title"]',
    )
  })
})