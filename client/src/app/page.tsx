import { Header } from '@/components/header'
import { NewTaskForm } from '@/components/new-task-form'
import { TaskList } from '@/components/task/task-list'

export default function Home() {
  return (
    <div className="flex min-h-screen items-start bg-background px-6 pb-8">
      <div className="mx-auto w-full max-w-2xl space-y-6">
        <Header />

        <NewTaskForm />

        <TaskList />
      </div>
    </div>
  )
}
