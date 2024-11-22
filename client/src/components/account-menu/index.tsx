'use client'

import { useQuery } from '@tanstack/react-query'
import { ChevronDown } from 'lucide-react'

import { getProfile, GetProfileResponse } from '@/http/get-profile'

import { Button } from '../ui/button'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from '../ui/dropdown-menu'
import { SignOutItem } from './sign-out-item'

export function AccountMenu() {
  const { data } = useQuery<GetProfileResponse>({
    queryKey: ['profile'],
    queryFn: getProfile,
  })

  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button
          variant="outline"
          className="items-cente group flex select-none gap-2 px-3"
        >
          {data?.username}
          <ChevronDown className="size-4" />
        </Button>
      </DropdownMenuTrigger>

      <DropdownMenuContent align="end" className="w-56">
        <DropdownMenuLabel>Menu</DropdownMenuLabel>

        <DropdownMenuSeparator />

        <SignOutItem />
      </DropdownMenuContent>
    </DropdownMenu>
  )
}
