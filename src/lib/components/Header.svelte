<script lang="ts">

    import Logo from '../assets/logo.png'

    import {page} from '$app/state'

</script>

<div class="flex gap-10 justify-around w-screen h-20 text-white bg-gray-900">
    <div class="flex items-center w-auto h-20">
        <img src={Logo} alt="Logo" class="rounded-3xl p-4 w-28 h-auto"/>
        <h1 class="flex self-center font-bold ">
            Juned Media
        </h1>
    </div>
    {#if page.url.pathname !== "/upload"}
        <a href="/upload" aria-label="upload" class="flex justify-center items-center">Upload</a>
    {/if}
    {#if page.url.pathname !== "/"}
        <a href="/" aria-label="upload" class="flex justify-center items-center">Home</a>
    {/if}
</div>