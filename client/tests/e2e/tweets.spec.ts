import { test } from '@playwright/test'
import { PlaywrightUser } from '../playwright-user'
import { USERS } from '../users'

test('user can tweet', async ({ page }) => {
  const playwrightUser = new PlaywrightUser(page)

  await playwrightUser.register(
    USERS.default.username,
    USERS.default.email,
    USERS.default.password
  )
  await playwrightUser.login(USERS.default.username, USERS.default.password)

  await playwrightUser.tweet('Hello, World!')
  await page.waitForSelector('text=Hello, World!')
})
