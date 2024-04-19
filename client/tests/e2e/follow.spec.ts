import { test, expect } from '@playwright/test'
import { PlaywrightUser } from '../playwright-user'
import { USERS } from '../users'

// Not finished
test('user A can follow user B', async ({ context }, testCase) => {
  const uniqueTestCase = testCase.testId
  const userA = `${uniqueTestCase}-A`
  const userB = `${uniqueTestCase}-B`

  console.log(uniqueTestCase)

  const page = await context.newPage()
  const test = new PlaywrightUser(page)

  // User B
  await test.register(userB, USERS.default.email, USERS.default.password)
  await test.login(userB, USERS.default.password)

  await test.tweet(`This is ${userB} tweet`)
  await test.logout()

  // User A
  await test.register(userA, USERS.default.email, USERS.default.password)
  await test.login(userA, USERS.default.password)

  await page.getByTestId('public-timeline-link').click()

  // Follow
  await page
    .locator('p')
    .filter({ hasText: `This is ${userB} tweet` })
    .getByRole('link')
    .first()
    .click()

  await page.getByTestId('follow').click()

  await page.getByTestId('my-timeline-link').click()

  await expect(page.locator('p').first()).toContainText(
    `This is ${userB} tweet`
  )

  // Unfollow
  await page
    .locator('p')
    .filter({ hasText: `This is ${userB} tweet` })
    .getByRole('link')
    .first()
    .click()

  await page.getByTestId('unfollow').click()

  await page.getByTestId('my-timeline-link').click()

  await expect(
    await page
      .locator('p')
      .filter({ hasText: `This is ${userB} tweet` })
      .getByRole('link')
  ).toHaveCount(0)
})
