import { TimelineSkeleton } from '@/components/TimelineSkeleton'
import { MessageList } from '@/components/MessageList'
import { PageWrapper } from '@/components/PageWrapper'
import { TweetBox } from '@/components/TweetBox'
import { getTimeline } from '@/services/api'
import { TimelineResponse } from '@/services/api.types'
import { useCallback, useEffect, useState } from 'react'
import { toast } from 'react-toastify'

export const MyTimeline = () => {
  const [timeline, setTimeline] = useState<TimelineResponse>()
  const [isLoading, setIsLoading] = useState(true)

  const fetchMessages = useCallback(async () => {
    try {
      setIsLoading(true)
      const response = await getTimeline()

      setTimeline(response.data)
    } catch (error) {
      console.error(error)
    } finally {
      setIsLoading(false)
    }
  }, [])

  useEffect(() => {
    fetchMessages()
  }, [fetchMessages])

  const handleAddMessageSuccess = () => {
    fetchMessages()
    toast.success('Your message was recorded')
  }

  return (
    <PageWrapper data-testid='my-timeline-page'>
      <h2>My Timeline</h2>

      <TimelineSkeleton isLoading={isLoading}>
        <>
          {timeline?.User && (
            <TweetBox
              handleCTA={handleAddMessageSuccess}
              username={timeline.User.Username}
            />
          )}

          <MessageList messages={timeline?.Messages} />
        </>
      </TimelineSkeleton>
    </PageWrapper>
  )
}
