import { useEffect, useState } from 'react'
import { User } from '@/services/api.types'
import axios from 'axios'

type Props = {
  user: User | undefined
  profile: User | undefined
  isByDefaultFollowing: boolean | undefined
  setFlashMessage: (message: string) => void
}

export const FollowStatus = ({
  user,
  profile,
  isByDefaultFollowing = false,
  setFlashMessage,
}: Props) => {
  const [following, setFollowing] = useState<boolean>(false)
  const profileUsername = profile?.Username

  useEffect(() => {
    setFollowing(isByDefaultFollowing)
  }, [isByDefaultFollowing])

  const handleUnfollow = async () => {
    try {
      await axios.post(`/api/${profileUsername}/unfollow`)
      setFlashMessage(`You are no longer following "${profileUsername}"`)
      setFollowing(false)
    } catch (error) {
      console.error(error)
    }
  }

  const handleFollow = async () => {
    try {
      await axios.post(`/api/${profileUsername}/follow`)
      setFlashMessage(`You are now following "${profileUsername}"`)
      setFollowing(true)
    } catch (error) {
      console.error(error)
    }
  }

  const renderFollowStatus = () => {
    if (profile?.User_id === user?.User_id) {
      return 'This is you!'
    }

    if (following) {
      return (
        <>
          You are currently following this user.{' '}
          <button className='unfollow' onClick={handleUnfollow}>
            Unfollow user
          </button>
          .
        </>
      )
    }

    return (
      <>
        You are not yet following this user.{' '}
        <button className='unfollow' onClick={handleFollow}>
          Follow user
        </button>
        .
      </>
    )
  }

  if (!user || !profile) return null
  return <div className='followstatus'>{renderFollowStatus()}</div>
}
