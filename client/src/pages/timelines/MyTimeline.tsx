import { TimelineSkeleton } from '@/components/TimelineSkeleton'
import { MessageList } from '@/components/MessageList'
import { PageWrapper } from '@/components/PageWrapper'
import { TweetBox } from '@/components/TweetBox'
import { getTimeline } from '@/services/api'
import { TimelineResponse } from '@/services/api.types'
import { useCallback, useEffect, useState } from 'react'
import { useLocation } from 'react-router-dom'

export const MyTimeline = () => {
  const { state } = useLocation()

  const [timeline, setTimeline] = useState<TimelineResponse>()
  const [flashMessage, setFlashMessage] = useState<string | null>(null)
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
    setFlashMessage('Your message was recorded')
  }

  return (
    <PageWrapper
      data-testid='my-timeline-page'
      flashMessage={flashMessage || (state && state.flashMessage)}
    >
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
