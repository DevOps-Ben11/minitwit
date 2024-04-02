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

  const fetchMessages = useCallback(async () => {
    const response = await getTimeline()

    setTimeline(response.data)
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

      {timeline?.User && (
        <TweetBox
          handleCTA={handleAddMessageSuccess}
          username={timeline.User.Username}
        />
      )}

      <MessageList messages={timeline?.Messages} />
    </PageWrapper>
  )
}
