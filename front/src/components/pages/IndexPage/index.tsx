import React, { useState, useCallback } from 'react'
import { filter } from 'graphql-anywhere'
import {
  useIndexPageQuery,
  RoomListFragment,
  RoomListFragmentDoc,
} from '@/graphql'
import { useAuthContext } from '@/contexts/auth'
import Skeleton from './Skeleton'
import Component from './presenter'

const IndexPageContainer: React.VFC = () => {
  const { isLoggedIn } = useAuthContext()
  const [openModal, setOpenModal] = useState(false)
  const { data } = useIndexPageQuery({ fetchPolicy: 'network-only' })

  const handleOpenModal = useCallback(() => {
    setOpenModal(true)
  }, [])

  const handleCloseModal = useCallback(() => {
    setOpenModal(false)
  }, [])

  const rooms =
    data && filter<RoomListFragment>(RoomListFragmentDoc, data).rooms.nodes

  if (!rooms) return <Skeleton />

  return (
    <Component
      contentHeaderProps={{ isLoggedIn, handleOpenModal }}
      roomListProps={{ rooms }}
      createRoomModalProps={{ open: openModal, handleClose: handleCloseModal }}
    />
  )
}

export default IndexPageContainer
