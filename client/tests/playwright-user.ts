import { type Page } from '@playwright/test'

export class PlaywrightUser {
  readonly page: Page

  constructor(page: Page) {
    this.page = page
  }

  async register(username: string, email: string, password: string) {
    await this.page.goto('/register')

    await this.page.fill('input[name=username]', username)
    await this.page.fill('input[name=email]', email)
    await this.page.fill('input[name=password]', password)
    await this.page.fill('input[name=passwordRepeat]', password)
    await this.page.click('input[type=submit]')

    await this.page.waitForTimeout(500)
  }

  async login(username: string, password: string) {
    await this.page.goto('/login')

    await this.page.fill('input[name=username]', username)
    await this.page.fill('input[name=password]', password)
    await this.page.click('input[type=submit]')

    // Wait for page redirect
    await this.page.waitForURL('**/', { timeout: 5000 })
  }

  async tweet(tweet: string) {
    await this.page.fill('input[name=message]', tweet)
    await this.page.click('input[type=submit]')
  }

  async logout() {
    await this.page.getByTestId('sign-out').click()

    // Wait for page redirect
    await this.page.waitForURL('**/', { timeout: 5000 })
  }
}
