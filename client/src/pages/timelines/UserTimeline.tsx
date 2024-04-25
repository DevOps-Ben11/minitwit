import { FollowStatus } from '@/components/FollowStatus'
import { TimelineSkeleton } from '@/components/TimelineSkeleton'
import { MessageList } from '@/components/MessageList'
import { PageWrapper } from '@/components/PageWrapper'
import { getUserTimeline } from '@/services/api'
import { UserTimelineResponse } from '@/services/api.types'
import { useEffect, useState } from 'react'
import { useNavigate, useParams } from 'react-router-dom'

export const UserTimeline = () => {
  const { username } = useParams<{ username: string }>()
  const navigate = useNavigate()

  const [timeline, setTimeline] = useState<UserTimelineResponse>()
  const [isLoading, setIsLoading] = useState(true)

  useEffect(() => {
    if (!username) {
      return
    }

    const fetchMessages = async () => {
      try {
        const response = await getUserTimeline(username)
        setTimeline(response.data)
      } catch (error) {
        navigate('/public')
      } finally {
        setIsLoading(false)
      }
    }

    fetchMessages()
  }, [username, navigate])

  return (
    <PageWrapper>
      <h2>{username}'s Timeline </h2>

      <TimelineSkeleton isLoading={isLoading}>
        <>
          <FollowStatus
            user={timeline?.User}
            profile={timeline?.Profile}
            isByDefaultFollowing={timeline?.Followed}
          />

          <MessageList messages={timeline?.Messages} />
        </>
      </TimelineSkeleton>
    </PageWrapper>
  )
}
