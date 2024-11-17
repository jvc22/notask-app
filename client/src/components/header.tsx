import { GitHubLogoIcon } from '@radix-ui/react-icons'

import { ThemeToggle } from './theme/theme-toggle'
import { Button } from './ui/button'

export function Header() {
  return (
    <header className="flex w-full items-center justify-between py-4">
      <h1 className="text-lg font-bold">
        notask<span className="text-muted-foreground">.app</span>
      </h1>

      <div className="text-foreground ml-auto flex items-center gap-2 text-sm">
        <div>
          <Button size="sm" variant="ghost">
            <GitHubLogoIcon />
          </Button>

          <ThemeToggle />
        </div>
      </div>
    </header>
  )
}
