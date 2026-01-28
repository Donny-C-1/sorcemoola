<script>
	import { enhance } from "$app/forms";
    import Button from "$lib/components/ui/Button.svelte";

    let { form } = $props();
</script>

<h1>Create Campaign</h1>
<form action="?/createCampaign" method="post" use:enhance>
    <div>
        <label for="campaign_name">Campaign Name:</label>
        <input type="text" name="campaign_name" id="campaign_name" />
    </div>
    <br/>
    <div>
        <label for="campaign_story">Campaign Story:</label>
        <textarea name="campaign_story" id="campaign_story" rows="5"></textarea>
    </div>
    <br/>
    <div>
        <label for="fund_goal">Fund Goal (₦):</label>
        <input type="number" id="fund_goal" name="fund_goal" />
    </div>
    <div>
        <p>{form?.error?.label}</p>
        <p>{form?.error?.message}</p>
    </div>
    <br/>
    <Button type="submit">Submit</Button>
</form>

<style>
    input, textarea {
        display: block;
        width: 100%;
    }
</style>