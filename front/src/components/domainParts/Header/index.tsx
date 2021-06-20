import React from 'react'
import Header from './presentation'
import { useAuthContext } from '@/contexts/auth'

const HeaderContainer: React.FC = () => {
  const { isLoggedIn } = useAuthContext()

  return <Header isLoggedIn={isLoggedIn} />
}

export default HeaderContainer