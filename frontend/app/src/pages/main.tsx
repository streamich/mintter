import {classnames} from '@app/utils/classnames'
import {lazy} from 'react'
import {ErrorBoundary, FallbackProps} from 'react-error-boundary'
import {Redirect} from 'wouter'
import {Route, Switch, useRoute} from '../components/router'
import '../styles/main.scss'

var PublicationList = lazy(() => import('@app/pages/publication-list-page'))
var DraftList = lazy(() => import('@app/pages/draft-list-page'))
var Publication = lazy(() => import('@app/pages/publication'))
var Draft = lazy(() => import('@app/pages/draft'))
var Settings = lazy(() => import('@app/pages/settings'))
var QuickSwitcher = lazy(() => import('@components/quick-switcher'))
var Topbar = lazy(() => import('@components/topbar'))
var Footer = lazy(() => import('@components/footer'))

export default function Main() {
  let [isSettings] = useRoute('/settings')
  return (
    <ErrorBoundary
      FallbackComponent={MainBoundary}
      onReset={() => {
        window.location.reload()
      }}
    >
      <div className={classnames('main-root', {settings: isSettings})}>
        <main>
          <Switch>
            {/* <Route path="/" component={Home} /> */}
            <Route path="/inbox" component={PublicationList} />
            <Route path="/drafts" component={DraftList} />
            <Route path="/p/:id/:version/:block?" component={Publication} />
            <Route path="/d/:id" component={Draft} />
            <Route path="/settings" component={Settings} />
            <Route>{() => <Redirect to="/inbox" />}</Route>
          </Switch>
        </main>
        {!isSettings ? (
          <>
            <Topbar />
            <QuickSwitcher />
            <Footer />
          </>
        ) : null}
      </div>
    </ErrorBoundary>
  )
}

function MainBoundary({error, resetErrorBoundary}: FallbackProps) {
  return (
    <div role="alert" data-layout-section="main">
      <p>Publication Error</p>
      <pre>{error.message}</pre>
      <button onClick={resetErrorBoundary}>reload page</button>
    </div>
  )
}