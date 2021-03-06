import React, { createContext, useContext, useState } from 'react'
import { GlobalUser as CurrentUser } from '@/types/user'

type AuthContextValue = {
  // undefined : まだログイン確認が完了していない状態とする
  // null      : ログイン確認をした結果、ログインしていなかった状態とする
  currentUser: undefined | null | CurrentUser
  setCurrentUser: (user: CurrentUser | null) => void
}

const defaultContextValue: AuthContextValue = {
  currentUser: undefined,
  setCurrentUser: () => {},
}

export const AuthContext = createContext<AuthContextValue>(defaultContextValue)
export const useAuthContext = () => {
  const { currentUser, ...props } = useContext(AuthContext)
  const isAuthChecking = currentUser === undefined
  const isLoggedIn = !!currentUser

  return {
    currentUser,
    isAuthChecking,
    isLoggedIn,
    ...props,
  }
}
export const useCurrentUser = () => {
  const { currentUser } = useContext(AuthContext)
  const isAuthChecking = currentUser === undefined

  return {
    currentUser,
    isAuthChecking,
  }
}

export const AuthProvider: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  const [currentUser, setCurrentUser] = useState<
    undefined | null | CurrentUser
  >(undefined)

  const values: AuthContextValue = {
    currentUser,
    setCurrentUser,
  }

  return <AuthContext.Provider value={values}>{children}</AuthContext.Provider>
}
