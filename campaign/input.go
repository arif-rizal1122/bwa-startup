package campaign

type GetCampaingDetailInput struct {
	ID            int       `uri:"id" binding:"required"`
}

