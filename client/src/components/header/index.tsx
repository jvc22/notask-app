import { AccountMenu } from '../account-menu'
import { ThemeToggle } from '../theme/theme-toggle'
import { GitHubButton } from './github-button'

export function Header() {
  return (
    <header className="flex w-full items-center justify-between pt-6">
      <h1 className="text-lg font-bold">
        notask<span className="text-muted-foreground">.app</span>
      </h1>

      <div className="ml-auto flex items-center gap-2 text-sm text-foreground">
        <div>
          <GitHubButton />

          <ThemeToggle />
        </div>
        <AccountMenu />
      </div>
    </header>
  )
}
