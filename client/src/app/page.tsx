import { Header } from '@/components/header'
import { NewTaaskForm } from '@/components/new-task-form'

export default function Home() {
  return (
    <div className="bg-background flex min-h-screen items-start px-6 pb-8">
      <div className="mx-auto w-full max-w-2xl space-y-6">
        <Header />

        <NewTaaskForm />
      </div>
    </div>
  )
}
