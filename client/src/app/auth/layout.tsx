import { Metadata } from 'next'
import { redirect } from 'next/navigation'

import { isAuthenticated } from '@/auth/auth'

export const metadata: Metadata = {
  title: 'Auth | notask',
}

export default async function AuthLayout({
  children,
}: Readonly<{
  children: React.ReactNode
}>) {
  if (await isAuthenticated()) {
    redirect('/')
  }

  return (
    <div className="flex min-h-screen flex-col items-center justify-center">
      <div className="w-full max-w-xs">{children}</div>
    </div>
  )
}
