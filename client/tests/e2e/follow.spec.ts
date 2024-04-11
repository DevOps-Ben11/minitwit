import { test, expect } from '@playwright/test'
import { PlaywrightUser } from '../playwright-user'
import { USERS } from '../users'

// Not finished
test('user A can follow user B', async ({ context }) => {
  const pageA = await context.newPage()
  const pageB = await context.newPage()

  const playwrightUserA = new PlaywrightUser(pageA)
  const playwrightUserB = new PlaywrightUser(pageB)

  // Login user A
  await playwrightUserA.register(
    USERS.a.username,
    USERS.a.email,
    USERS.a.password
  )
  await playwrightUserA.login(USERS.a.username, USERS.a.password)

  // Login user B
  await playwrightUserB.register(
    USERS.b.username,
    USERS.b.email,
    USERS.b.password
  )
  await playwrightUserB.login(USERS.b.username, USERS.b.password)

  await expect(pageB.getByTestId('my-timeline-page')).toBeVisible()
})
