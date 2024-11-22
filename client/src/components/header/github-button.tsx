'use client'

import { GitHubLogoIcon } from '@radix-ui/react-icons'

import { Button } from '../ui/button'

export function GitHubButton() {
  function handleOpenGitHub() {
    window.open('https://github.com/jvc22/notask-app', '_blank')
  }

  return (
    <Button size="sm" variant="ghost" onClick={handleOpenGitHub}>
      <GitHubLogoIcon />
    </Button>
  )
}
