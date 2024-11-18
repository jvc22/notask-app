'use client'

import { QueryClientProvider } from '@tanstack/react-query'
import { ReactNode } from 'react'

import { getQueryClient } from '@/lib/react-query'

import { ThemeProvider } from '../components/theme/theme-provider'
import { Toaster } from '../components/ui/sonner'

export function Providers({ children }: { children: ReactNode }) {
  const queryClient = getQueryClient()

  return (
    <QueryClientProvider client={queryClient}>
      <ThemeProvider
        attribute="class"
        defaultTheme="dark"
        storageKey="notask-theme"
        enableSystem
      >
        <Toaster />
        {children}
      </ThemeProvider>
    </QueryClientProvider>
  )
}
