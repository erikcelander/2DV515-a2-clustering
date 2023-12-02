import Clusters from '@/components/clusters'
import { Button } from '@/components/ui/button'

export type ClusterData = {
  cluster1: string[];
  cluster2: string[];
  cluster3: string[];
  cluster4: string[];
  cluster5: string[];
}


async function getClusters(): Promise<ClusterData> {
  const res = await fetch('http://localhost:8080/clusters');
  const data: ClusterData = await res.json();
  return data;
}



export default async function Home() {
  const clusters = await getClusters();


  return (
    <main className='flex flex-col items-center justify-between p-8'>
      <Button className='mb-8'>Fetch</Button>
      <Clusters clusters={clusters} />
    </main>
  )
}
