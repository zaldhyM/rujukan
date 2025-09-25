<script lang="ts">
  import { Button, Modal } from "@sveltestrap/sveltestrap";
  let open = false;
  const toggle = () => (open = !open);
</script>

<!-- Page Heading -->
<div class="d-sm-flex align-items-center justify-content-between mb-4">
  <h1 class="h3 mb-0 text-gray-800">Dashbord</h1>
  <Button color="danger" on:click={toggle}>Open Modal</Button>
  <Modal body header="Modal title" isOpen={open} {toggle}>
    Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod
    tempor incididunt ut labore et dolore magna aliqua.
  </Modal>
</div>
