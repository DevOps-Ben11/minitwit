import { test, expect } from '@playwright/test'

test('can see timeline', async ({ page }) => {
  // Register user
  await page.goto('/register')

  // Fill the form.
  await page.fill('input[name=username]', 'playwrightUser')
  await page.fill('input[name=email]', 'playwrightUser@gmail.com')
  await page.fill('input[name=password]', 'password1')
  await page.fill('input[name=passwordRepeat]', 'password1')
  await page.click('input[type=submit]')

  // Login user
  await page.waitForURL('**/login', { timeout: 5000 })

  // Fill the form.
  await page.fill('input[name=username]', 'playwrightUser')
  await page.fill('input[name=password]', 'password1')
  await page.click('input[type=submit]')

  await page.waitForURL('**/', { timeout: 5000 })

  await expect(page.getByTestId('my-timeline-page')).toBeVisible()
})
