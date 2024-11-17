import { ReactNode } from 'react'

import { ThemeProvider } from '@/components/theme/theme-provider'
import { Toaster } from '@/components/ui/sonner'

export function Providers({ children }: { children: ReactNode }) {
  return (
    <ThemeProvider
      attribute="class"
      defaultTheme="dark"
      storageKey="notask-theme"
      enableSystem
    >
      <Toaster />
      {children}
    </ThemeProvider>
  )
}
