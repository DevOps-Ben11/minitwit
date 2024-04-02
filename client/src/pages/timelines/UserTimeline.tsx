import { MessageList } from '@/components/MessageList'
import { PageWrapper } from '@/components/PageWrapper'
import { getUserTimeline } from '@/services/api'
import { UserTimelineResponse } from '@/services/api.types'
import axios from 'axios'
import { useEffect, useState } from 'react'
import { useNavigate, useParams } from 'react-router-dom'

export const Timeline = () => {
  const { username } = useParams<{ username: string }>()
  const navigate = useNavigate()

  const [timeline, setTimeline] = useState<UserTimelineResponse>()
  const [following, setFollowing] = useState<boolean>(false)
  const [flashMessage, setFlashMessage] = useState<string | null>(null)

  useEffect(() => {
    if (!username) {
      return
    }

    const fetchMessages = async () => {
      try {
        const response = await getUserTimeline(username)

        setTimeline(response.data)
        setFollowing(response.data.Followed)
      } catch (error) {
        navigate('/')
      }
    }

    fetchMessages()
  }, [username, navigate])

  const handleUnfollow = async () => {
    try {
      await axios.post(`/api/${username}/unfollow`)
      setFlashMessage(`You are no longer following "${username}"`)
      setFollowing(false)
    } catch (error) {
      console.error(error)
    }
  }

  const handleFollow = async () => {
    try {
      await axios.post(`/api/${username}/follow`)
      setFlashMessage(`You are now following "${username}"`)
      setFollowing(true)
    } catch (error) {
      console.error(error)
    }
  }

  const renderFollowStatus = () => {
    if (timeline?.Profile?.User_id === timeline?.User?.User_id) {
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

  return (
    <PageWrapper flashMessage={flashMessage}>
      <h2>{username}'s Timeline </h2>

      {timeline?.User && (
        <div className='followstatus'>{renderFollowStatus()}</div>
      )}

      <MessageList messages={timeline?.Messages} />
    </PageWrapper>
  )
}
