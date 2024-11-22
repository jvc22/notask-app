import { deleteCookie } from 'cookies-next'
import { LogOut } from 'lucide-react'
import { useRouter } from 'next/navigation'

import { DropdownMenuItem } from '../ui/dropdown-menu'

export function SignOutItem() {
  const router = useRouter()

  function handleSignOut() {
    deleteCookie('notask-token')

    router.push('/auth/sign-in')
  }

  return (
    <DropdownMenuItem
      onClick={handleSignOut}
      className="text-rose-500 dark:text-rose-400"
    >
      <LogOut className="mr-2 size-4" />

      <span>Log out</span>
    </DropdownMenuItem>
  )
}
