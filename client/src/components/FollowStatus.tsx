import { useEffect, useState } from 'react'
import { User } from '@/services/api.types'
import { followUser, unFollowUser } from '@/services/api'
import { toast } from 'react-toastify'

type Props = {
  user: User | undefined
  profile: User | undefined
  isByDefaultFollowing: boolean | undefined
}

export const FollowStatus = ({
  user,
  profile,
  isByDefaultFollowing = false,
}: Props) => {
  const [following, setFollowing] = useState<boolean>(false)

  useEffect(() => {
    setFollowing(isByDefaultFollowing)
  }, [isByDefaultFollowing])

  const handleUnFollow = async (username: string) => {
    try {
      await unFollowUser(username)
      toast.success(`You are no longer following "${username}"`)
      setFollowing(false)
    } catch (error) {
      console.error(error)
    }
  }

  const handleFollow = async (username: string) => {
    try {
      await followUser(username)
      toast.success(`You are now following "${username}"`)
      setFollowing(true)
    } catch (error) {
      console.error(error)
    }
  }

  const renderFollowStatus = (username: string) => {
    if (profile?.User_id === user?.User_id) {
      return 'This is you!'
    }

    if (following) {
      return (
        <>
          You are currently following this user.{' '}
          <button
            data-testid='unfollow'
            className='unfollow'
            onClick={() => handleUnFollow(username)}
          >
            Unfollow user
          </button>
          .
        </>
      )
    }

    return (
      <>
        You are not yet following this user.{' '}
        <button
          data-testid='follow'
          className='follow'
          onClick={() => handleFollow(username)}
        >
          Follow user
        </button>
        .
      </>
    )
  }

  if (!user || !profile) return null
  return (
    <div className='followstatus'>{renderFollowStatus(profile.Username)}</div>
  )
}
