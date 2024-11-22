import { Metadata } from 'next'
import { redirect } from 'next/navigation'

import { isAuthenticated } from '@/auth/auth'

export const metadata: Metadata = {
  title: 'Home | notask',
}

export default async function AppLayout({
  children,
}: Readonly<{
  children: React.ReactNode
}>) {
  if (!(await isAuthenticated())) {
    redirect('/auth/sign-in')
  }

  return (
    <div className="flex min-h-screen items-start bg-background">
      {children}
    </div>
  )
}
