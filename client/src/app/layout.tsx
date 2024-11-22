import './globals.css'

import type { Metadata } from 'next'
import localFont from 'next/font/local'

import { Providers } from './providers'

const geistSans = localFont({
  src: './fonts/GeistVF.woff',
  variable: '--font-geist-sans',
  weight: '100 900',
})
const geistMono = localFont({
  src: './fonts/GeistMonoVF.woff',
  variable: '--font-geist-mono',
  weight: '100 900',
})

export const metadata: Metadata = {
  title: 'App | notask',
  description: 'Minimal To-Do List app',
}

export default async function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode
}>) {
  return (
    <html lang="en">
      <body
        className={`${geistSans.variable} ${geistMono.variable} px-6 font-geist antialiased`}
      >
        <Providers>{children}</Providers>
      </body>
    </html>
  )
}
