import {InterpreterFrom} from 'xstate'
import {createAuthService} from './auth-machine'
import {createInterpreterContext} from './utils/machine-utils'
const [AuthProvider, useAuthService, createAuthSelector] =
  createInterpreterContext<
    InterpreterFrom<ReturnType<typeof createAuthService>>
  >('Auth')

export {AuthProvider, useAuthService}

export const useAccountInfo = createAuthSelector(
  (state) => state.context.accountInfo,
)

export const useAccount = createAuthSelector((state) => state.context.account)

export const useAccountProfile = createAuthSelector(
  (state) => state.context.account?.profile,
)

export const usePeerAddrs = createAuthSelector(
  (state) => state.context.peerAddrs,
)
