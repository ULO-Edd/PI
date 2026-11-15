import { test, expect } from "@playwright/test";

test.describe("Frontend Event Catalogue Accessibility and Navigation Checks", () => {
  test.beforeEach(async ({ page }) => {
    // Intercept API events request to provide reliable mock data
    await page.route("**/api/events", async (route) => {
      await route.fulfill({
        status: 200,
        contentType: "application/json",
        body: JSON.stringify([
          {
            id: "11111111-2222-3333-4444-555555555555",
            title: "Tech Conference 2026",
            description: "An incredible gather-up of tech innovators.",
            dateTime: "2026-10-15T09:00:00Z",
            location: "San Francisco, CA",
            capacity: 500,
            availableTickets: 350,
          },
          {
            id: "22222222-3333-4444-5555-666666666666",
            title: "Vue 3 Deep Dive",
            description: "In-depth workshop on reactive patterns and Pinia store mechanics.",
            dateTime: "2026-11-20T14:00:00Z",
            location: "Online Workshop",
            capacity: 100,
            availableTickets: 0, // Sold out
          },
        ]),
      });
    });

    // Navigate to local server landing page
    await page.goto("http://localhost:5173/");
  });

  test("should load the page with proper semantic landmarks and headings", async ({ page }) => {
    // Assert page title
    await expect(page).toHaveTitle(/SDD-Project3/);

    // Verify main landmarks exist for screen readers
    const main = page.locator("main");
    await expect(main).toBeVisible();

    const bannerSection = page.locator('section[aria-labelledby="hero-title"]');
    await expect(bannerSection).toBeVisible();

    const catalogueSection = page.locator('section[aria-label="Event Catalogue"]');
    await expect(catalogueSection).toBeVisible();

    // Verify heading sequence starts from high-level
    const h1 = page.locator("h1");
    await expect(h1).toHaveText("Seamless Event Discovery");

    const h3Elements = page.locator("h3");
    await expect(h3Elements.first()).toHaveText("Tech Conference 2026");
  });

  test("should provide accessible search control and live-region status reports", async ({ page }) => {
    const searchInput = page.locator("#search-input");
    await expect(searchInput).toBeVisible();
    await expect(searchInput).toHaveAttribute("placeholder", "Filter by title or description...");

    // Search query input and filtering
    await searchInput.fill("Vue 3");

    // Live status update announcement
    const statusRegion = page.locator('[role="status"]');
    await expect(statusRegion).toContainText("Showing 1 of 2 events");

    // Only "Vue 3 Deep Dive" event card should remain visible
    const eventTitle = page.locator("#event-title");
    await expect(eventTitle).toHaveText("Vue 3 Deep Dive");
  });
});
