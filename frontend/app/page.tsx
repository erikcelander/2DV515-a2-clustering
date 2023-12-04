import KmeansClusters from '@/components/kmeans-clusters'
import { Button } from '@/components/ui/button'

export type KmeansCluster = string[]

export type KmeansData = {
  cluster1: KmeansCluster
  cluster2: KmeansCluster
  cluster3: KmeansCluster
  cluster4: KmeansCluster
  cluster5: KmeansCluster
}

export type HierarchicalNode = {
  name?: string
  children?: HierarchicalNode[]
}

export type HierarchicalData = HierarchicalNode[]

export type ClusterData = {
  kmeans: KmeansData
  hierarchical: HierarchicalData
}

async function getClusters(): Promise<ClusterData> {
  const res = await fetch('http://localhost:8080/api/clusters', {
    cache: 'no-cache',
  })
  const data: ClusterData = await res.json()
  return data
}

export default async function Home() {
  const { kmeans, hierarchical } = await getClusters()

  return (
    <main className='flex flex-col items-center justify-between p-8'>
      <Button className='mb-8'>Fetch</Button>
      {/*<Clusters clusters={clusters} />*/}
    </main>
  )
}
