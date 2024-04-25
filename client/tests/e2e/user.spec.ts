import { test, expect } from '@playwright/test'
import { PlaywrightUser } from '../playwright-user'
import { USERS } from '../users'

test('can register and login a user', async ({ page }) => {
  const playwrightUser = new PlaywrightUser(page)

  await playwrightUser.register(
    USERS.default.username,
    USERS.default.email,
    USERS.default.password
  )
  await playwrightUser.login(USERS.default.username, USERS.default.password)

  await expect(page.getByTestId('my-timeline-page')).toBeVisible()
})

test('user cannot access my timeline if not logged in', async ({ page }) => {
  await page.goto('/')

  await expect(page.getByTestId('public-timeline-page')).toBeVisible()
})
