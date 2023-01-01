
# Analytics V1 to V2 Migration

Using your existing Notion database:
1. Create an property of type `Number` named `Interactions`.
2. Create a property of type `Select` named `Interact` with the options `Enabled`, `Disabled`, and `Dynamic`.
3. Ensure you have a property called `CreatedBy` of type `Created by`.
4. Enable your database's `Subtasks` feature (using the `ParentCampaign` and `SubCampaigns` property names respectively) for better visual organization. Alternatively, just ensure the `ParentCampaign` of type `Relation` property exists.

Then, you can deploy your new Analytics instance (example instructions in README) and start using the new APIs, documented there as well.

Your new Analytics instance and database format is still compatible with your `redirect` instance, if you are running the two alongside each other.