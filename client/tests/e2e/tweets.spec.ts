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

  await page.fill('input[name=message]', 'Hello, World!')
  await page.click('input[type=submit]')
  await page.waitForSelector('text=Hello, World!')
})
