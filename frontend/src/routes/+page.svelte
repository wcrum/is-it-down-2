<script lang="ts">
  import Activity from "lucide-svelte/icons/activity";
  import ArrowUpRight from "lucide-svelte/icons/arrow-up-right";
  import CircleUser from "lucide-svelte/icons/circle-user";
  import CreditCard from "lucide-svelte/icons/credit-card";
  import DollarSign from "lucide-svelte/icons/dollar-sign";
  import Menu from "lucide-svelte/icons/menu";
  import Package2 from "lucide-svelte/icons/package-2";
  import Search from "lucide-svelte/icons/search";
  import Users from "lucide-svelte/icons/users";
  

  import Clipboard from "lucide-svelte/icons/clipboard-check"
  import Network from "lucide-svelte/icons/network"

  import * as Avatar from "$lib/components/ui/avatar/index.js";
  import { Badge } from "$lib/components/ui/badge/index.js";
  import { Button } from "$lib/components/ui/button/index.js";
  import * as Card from "$lib/components/ui/card/index.js";
  import * as DropdownMenu from "$lib/components/ui/dropdown-menu/index.js";
  import { Input } from "$lib/components/ui/input/index.js";
  import * as Sheet from "$lib/components/ui/sheet/index.js";
  import * as Table from "$lib/components/ui/table/index.js";
  import Chart from "chart.js/auto";

  let dataToGraph = [0, 100, 100, 100, 100, 100, 100, 100];
  let chartObject;

  function chart(node, data) {
    function setupChart(_data) {
      chartObject = new Chart(node, {
        type: "line",
        data: {
          labels: [
            "Red",
            "Blue",
            "Yellow",
            "Green",
            "Purple",
            "Orange",
            "Purple2",
            "Val2",
            "Graph2",
            "Dap",
          ],
          datasets: [
            {
              label: "# of Votes",
              data: _data,
              borderWidth: 1,
            },
          ],
        },
        options: {
          responsive: true,
          maintainAspectRatio: false,
          scales: {
            y: {
              beginAtZero: true,
              display: false,
            },
            x: {
              display: false,
            },
          },
          plugins: {
            legend: {
              display: false,
            },
          },
        },
      });
    }
    setupChart(data);
    return {
      update(data) {
        chartObject.destroy();
        setupChart(data);
      },
      destroy() {
        chartObject.destroy();
      },
    };
  }
</script>
<div class="flex min-h-screen w-full flex-col dark">
  <header
    class="bg-background sticky top-0 flex h-16 items-center gap-4 border-b px-4 md:px-6"
  >
    <nav
      class="hidden flex-col gap-6 text-lg font-medium md:flex md:flex-row md:items-center md:gap-5 md:text-sm lg:gap-6"
    >
      <a
        href="##"
        class="flex items-center gap-2 text-lg font-semibold md:text-base"
      >
        <Package2 class="h-6 w-6" />
        <span class="sr-only">Acme Inc</span>
      </a>
      <a
        href="##"
        class="text-foreground hover:text-foreground transition-colors"
      >
        Dashboard
      </a>
      <a
        href="##"
        class="text-muted-foreground hover:text-foreground transition-colors"
      >
        Orders
      </a>
      <a
        href="##"
        class="text-muted-foreground hover:text-foreground transition-colors"
      >
        Products
      </a>
      <a
        href="##"
        class="text-muted-foreground hover:text-foreground transition-colors"
      >
        Customers
      </a>
      <a
        href="##"
        class="text-muted-foreground hover:text-foreground transition-colors"
      >
        Analytics
      </a>
    </nav>
    <Sheet.Root>
      <Sheet.Trigger asChild let:builder>
        <Button
          variant="outline"
          size="icon"
          class="shrink-0 md:hidden"
          builders={[builder]}
        >
          <Menu class="h-5 w-5" />
          <span class="sr-only">Toggle navigation menu</span>
        </Button>
      </Sheet.Trigger>
      <Sheet.Content side="left">
        <nav class="grid gap-6 text-lg font-medium">
          <a href="##" class="flex items-center gap-2 text-lg font-semibold">
            <Package2 class="h-6 w-6" />
            <span class="sr-only">Acme Inc</span>
          </a>
          <a href="##" class="hover:text-foreground"> Dashboard </a>
          <a href="##" class="text-muted-foreground hover:text-foreground">
            Orders
          </a>
          <a href="##" class="text-muted-foreground hover:text-foreground">
            Products
          </a>
          <a href="##" class="text-muted-foreground hover:text-foreground">
            Customers
          </a>
          <a href="##" class="text-muted-foreground hover:text-foreground">
            Analytics
          </a>
        </nav>
      </Sheet.Content>
    </Sheet.Root>
    <div class="flex w-full items-center gap-4 md:ml-auto md:gap-2 lg:gap-4">
      <form class="ml-auto flex-1 sm:flex-initial">
        <div class="relative">
          <Search
            class="text-muted-foreground absolute left-2.5 top-2.5 h-4 w-4"
          />
          <Input
            type="search"
            placeholder="Search products..."
            class="pl-8 sm:w-[300px] md:w-[200px] lg:w-[300px]"
          />
        </div>
      </form>
      <DropdownMenu.Root>
        <DropdownMenu.Trigger asChild let:builder>
          <Button
            builders={[builder]}
            variant="secondary"
            size="icon"
            class="rounded-full"
          >
            <CircleUser class="h-5 w-5" />
            <span class="sr-only">Toggle user menu</span>
          </Button>
        </DropdownMenu.Trigger>
        <DropdownMenu.Content align="end">
          <DropdownMenu.Label>My Account</DropdownMenu.Label>
          <DropdownMenu.Separator />
          <DropdownMenu.Item>Settings</DropdownMenu.Item>
          <DropdownMenu.Item>Support</DropdownMenu.Item>
          <DropdownMenu.Separator />
          <DropdownMenu.Item>Logout</DropdownMenu.Item>
        </DropdownMenu.Content>
      </DropdownMenu.Root>
    </div>
  </header>
  <main class="flex flex-1 flex-col gap-4 p-4 md:gap-8 md:p-8">
    <div class="grid gap-4 md:grid-cols-2 md:gap-8 lg:grid-cols-4">
      <Card.Root>
        <Card.Header
          class="flex flex-row items-center justify-between space-y-0 pb-2"
        >
          <Card.Title class="text-sm font-medium">Total Revenue</Card.Title>
          <DollarSign class="text-muted-foreground h-4 w-4" />
        </Card.Header>
        <Card.Content>
          <div class="text-2xl font-bold">$45,231.89</div>
          <p class="text-muted-foreground text-xs">+20.1% from last month</p>
        </Card.Content>
      </Card.Root>
      <Card.Root>
        <Card.Header
          class="flex flex-row items-center justify-between space-y-0 pb-2"
        >
          <Card.Title class="text-sm font-medium">Uptime</Card.Title>
          <Users class="text-muted-foreground h-4 w-4" />
        </Card.Header>
        <Card.Content>
          <div class="text-2xl font-bold">+2350</div>
          <p class="text-muted-foreground text-xs">+180.1% from last month</p>
        </Card.Content>
      </Card.Root>
      <Card.Root>
        <Card.Header
          class="flex flex-row items-center justify-between space-y-0 pb-2"
        >
          <Card.Title class="text-sm font-medium">Total Jobs</Card.Title>
          <Clipboard class="text-muted-foreground h-4 w-4" />
        </Card.Header>
        <Card.Content>
          <div class="text-2xl font-bold">12,234</div>
          <p class="text-muted-foreground text-xs">+19% from last month</p>
        </Card.Content>
      </Card.Root>
      <Card.Root>
        <Card.Header
          class="flex flex-row items-center justify-between space-y-0 pb-2"
        >
          <Card.Title class="text-sm font-medium">Active Workers</Card.Title>
          <Network class="text-muted-foreground h-4 w-4" />
        </Card.Header>
        <Card.Content>
          <div class="text-2xl font-bold">5</div>
          <p class="text-muted-foreground text-xs">+201 since last hour</p>
        </Card.Content>
      </Card.Root>
    </div>
    <div class="grid gap-4 md:gap-8 lg:grid-cols-2 xl:grid-cols-3">
      <Card.Root class="xl:col-span-2">
        <Card.Header class="flex flex-row items-center">
          <div class="grid gap-2">
            <Card.Title>Transactions</Card.Title>
            <Card.Description
              >Recent transactions from your store.</Card.Description
            >
          </div>
          <Button href="##" size="sm" class="ml-auto gap-1">
            View All
            <ArrowUpRight class="h-4 w-4" />
          </Button>
        </Card.Header>
        <Card.Content>
          <Table.Root>
            <Table.Header>
              <Table.Row>
                <Table.Head>Endpoint</Table.Head>
                <Table.Head class="xl:table.-column hidden">Type</Table.Head>
                <Table.Head class="xl:table.-column hidden">Status</Table.Head>
                <Table.Head class="xl:table.-column hidden">Date</Table.Head>
                <Table.Head class="text-right">Latency</Table.Head>
              </Table.Row>
            </Table.Header>
            <Table.Body>
              {#each { length: 5 } as _, i}
                <Table.Row>
                  <Table.Cell>
                    <div class="font-medium">Google</div>
                    <div class="text-muted-foreground hidden text-sm md:inline">
                      10.10.10.10
                    </div>
                  </Table.Cell>
                  <Table.Cell class="xl:table.-column hidden">Refund</Table.Cell
                  >
                  <Table.Cell class="xl:table.-column hidden">
                    <Badge class="text-xs" variant="outline">Declined</Badge>
                  </Table.Cell>
                  <Table.Cell
                    class="md:table.-cell xl:table.-column hidden lg:hidden"
                  >
                    2023-06-24
                  </Table.Cell>
                  <Table.Cell class="flex justify-end">
                    <span class="inline-block align-bottom my-auto pe-5">152.1112</span>
                    <div class="">
                      <canvas
                        class="flex w-1/2 chart"
                        use:chart={dataToGraph}
                      />
                    </div>
                  </Table.Cell>
                </Table.Row>
              {/each}
            </Table.Body>
          </Table.Root>
        </Card.Content>
      </Card.Root>
      <Card.Root>
        <Card.Header>
          <Card.Title>Recent Sales</Card.Title>
        </Card.Header>
        <Card.Content class="grid gap-8">
          <div class="flex items-center gap-4">
            <Avatar.Root class="hidden h-9 w-9 sm:flex">
              <Avatar.Image src="/avatars/01.png" alt="Avatar" />
              <Avatar.Fallback>OM</Avatar.Fallback>
            </Avatar.Root>
            <div class="grid gap-1">
              <p class="text-sm font-medium leading-none">Olivia Martin</p>
              <p class="text-muted-foreground text-sm">
                olivia.martin@email.com
              </p>
            </div>
            <div class="ml-auto font-medium">+$1,999.00</div>
          </div>
          <div class="flex items-center gap-4">
            <Avatar.Root class="hidden h-9 w-9 sm:flex">
              <Avatar.Image src="/avatars/02.png" alt="Avatar" />
              <Avatar.Fallback>JL</Avatar.Fallback>
            </Avatar.Root>
            <div class="grid gap-1">
              <p class="text-sm font-medium leading-none">Jackson Lee</p>
              <p class="text-muted-foreground text-sm">jackson.lee@email.com</p>
            </div>
            <div class="ml-auto font-medium">+$39.00</div>
          </div>
          <div class="flex items-center gap-4">
            <Avatar.Root class="hidden h-9 w-9 sm:flex">
              <Avatar.Image src="/avatars/03.png" alt="Avatar" />
              <Avatar.Fallback>IN</Avatar.Fallback>
            </Avatar.Root>
            <div class="grid gap-1">
              <p class="text-sm font-medium leading-none">Isabella Nguyen</p>
              <p class="text-muted-foreground text-sm">
                isabella.nguyen@email.com
              </p>
            </div>
            <div class="ml-auto font-medium">+$299.00</div>
          </div>
          <div class="flex items-center gap-4">
            <Avatar.Root class="hidden h-9 w-9 sm:flex">
              <Avatar.Image src="/avatars/04.png" alt="Avatar" />
              <Avatar.Fallback>WK</Avatar.Fallback>
            </Avatar.Root>
            <div class="grid gap-1">
              <p class="text-sm font-medium leading-none">William Kim</p>
              <p class="text-muted-foreground text-sm">will@email.com</p>
            </div>
            <div class="ml-auto font-medium">+$99.00</div>
          </div>
          <div class="flex items-center gap-4">
            <Avatar.Root class="hidden h-9 w-9 sm:flex">
              <Avatar.Image src="/avatars/05.png" alt="Avatar" />
              <Avatar.Fallback>SD</Avatar.Fallback>
            </Avatar.Root>
            <div class="grid gap-1">
              <p class="text-sm font-medium leading-none">Sofia Davis</p>
              <p class="text-muted-foreground text-sm">sofia.davis@email.com</p>
            </div>
            <div class="ml-auto font-medium">+$39.00</div>
          </div>
        </Card.Content>
      </Card.Root>
    </div>
  </main>
</div>

<style>
  :root {
    --vis-font-family: Inter, Arial, "Helvetica Neue", Helvetica, sans-serif;
    --vis-color-main: #4d8cfd;
    --vis-color-main-light: #d0e0fe;
    --vis-color-main-dark: #3c588a;
    --vis-color-grey: #2a2a2a;
  }
</style>
