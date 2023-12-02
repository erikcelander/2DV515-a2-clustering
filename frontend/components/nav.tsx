
import { ModeToggle } from './mode-toggle'


export function Nav({
  className,
  ...props
}: React.HTMLAttributes<HTMLElement>) {
  return (
    <div className={`flex justify-between items-center h-16 px-4`} {...props}>
      <div className='flex-initial pr-1'>
        <h3 className="text-xl font-bold cursor-pointer">Erik Kroon Celander</h3>
      </div>
      <div className="flex-grow text-center">
        <span className="text-l font-bold">2DV515 - Web Intelligence - Assignment 2</span>
      </div>
      <div className='flex-initial pl-40'>
        <ModeToggle />
      </div>
    </div>
  )

}
