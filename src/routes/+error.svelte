<script lang="ts">

    import { page } from "$app/state"

</script>

<div class="flex flex-col items-center justify-center h-screen bg-slate-950 text-white">

    <h1 class="text-4xl font-bold text-center">
        Unexpected error!
    </h1>
    <h2 class="text-2xl font-bold text-center">
        I'm sorry :(
    </h2>
    <div class="flex flex-col m-20">
        <h2 class="text-center font-semibold">
            Error details
        </h2>
        <p>
            {page.status} {page.error?.message}
        </p>
    </div>

</div>