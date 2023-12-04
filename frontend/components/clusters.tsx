import {
  Tabs,
  TabsContent,
  TabsList,
  TabsTrigger,
} from "@/components/ui/tabs"

import { KmeansData, HierarchicalData } from '@/app/page'
import KmeansClusters from './kmeans-clusters'
import HierarchicalTree from './hierarchical-tree'


export function Clusters({kmeans, hierarchical}: {kmeans: KmeansData, hierarchical: HierarchicalData}) {
  return (
    <Tabs defaultValue="kmeans" className="w-[500px]">
      <TabsList className="grid w-full grid-cols-2">
        <TabsTrigger value="kmeans">K-Means</TabsTrigger>
        <TabsTrigger value="hierarchical">Hierarchical</TabsTrigger>
      </TabsList>
      <TabsContent value="kmeans">
        <KmeansClusters kmeans={kmeans} />
      </TabsContent>
      <TabsContent value="hierarchical">
        <HierarchicalTree hierarchical={hierarchical} />
      </TabsContent>
    </Tabs>
  )
}
