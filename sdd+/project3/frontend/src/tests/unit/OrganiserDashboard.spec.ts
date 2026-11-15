import { describe, it, expect } from "vitest";
import { mount } from "@vue/test-utils";
import OrganiserEventForm from "../../components/OrganiserEventForm.vue";

describe("OrganiserEventForm Validation Unit Tests", () => {
  it("should prevent submission and report errors on empty fields", async () => {
    const wrapper = mount(OrganiserEventForm, {
      props: {
        isEdit: false,
        event: null,
      },
    });

    // Attempt to submit empty form
    await wrapper.find("form").trigger("submit.prevent");

    // Assert that validation blocks save event emission
    expect(wrapper.emitted("save")).toBeUndefined();

    // Verify appropriate validation alerts
    const titleError = wrapper.find('[role="alert"]');
    expect(titleError.exists()).toBe(true);
  });

  it("should enforce the maximum limit of 100 characters on event titles", async () => {
    const wrapper = mount(OrganiserEventForm, {
      props: {
        isEdit: false,
        event: null,
      },
    });

    // Set title longer than 100 characters
    const longTitle = "A".repeat(101);
    const titleInput = wrapper.find("#title");
    await titleInput.setValue(longTitle);

    // Set other required fields to valid values
    const tomorrow = new Date();
    tomorrow.setDate(tomorrow.getDate() + 1);
    const formattedDate = tomorrow.toISOString().slice(0, 16);
    await wrapper.find("#dateTime").setValue(formattedDate);
    await wrapper.find("#capacity").setValue(50);

    // Attempt submission
    await wrapper.find("form").trigger("submit.prevent");

    // Verify submission blocked
    expect(wrapper.emitted("save")).toBeUndefined();

    // Assert that the title validation error matches length rule
    const alert = wrapper.find('[role="alert"]');
    expect(alert.text()).toContain("Title cannot exceed 100 characters");
  });

  it("should prevent a past event date from being configured on creation", async () => {
    const wrapper = mount(OrganiserEventForm, {
      props: {
        isEdit: false,
        event: null,
      },
    });

    await wrapper.find("#title").setValue("Valid Conference Title");

    // Set past event date
    const yesterday = new Date();
    yesterday.setDate(yesterday.getDate() - 1);
    const formattedDate = yesterday.toISOString().slice(0, 16);
    await wrapper.find("#dateTime").setValue(formattedDate);
    await wrapper.find("#capacity").setValue(100);

    // Submit form
    await wrapper.find("form").trigger("submit.prevent");

    expect(wrapper.emitted("save")).toBeUndefined();

    const alert = wrapper.find('[role="alert"]');
    expect(alert.text()).toContain("Event Date and Time must be in the future");
  });

  it("should validate that capacity is at least 1 registration", async () => {
    const wrapper = mount(OrganiserEventForm, {
      props: {
        isEdit: false,
        event: null,
      },
    });

    await wrapper.find("#title").setValue("Valid Conference Title");

    const tomorrow = new Date();
    tomorrow.setDate(tomorrow.getDate() + 1);
    const formattedDate = tomorrow.toISOString().slice(0, 16);
    await wrapper.find("#dateTime").setValue(formattedDate);

    // Set invalid capacity
    await wrapper.find("#capacity").setValue(0);

    // Submit form
    await wrapper.find("form").trigger("submit.prevent");

    expect(wrapper.emitted("save")).toBeUndefined();

    const alert = wrapper.find('[role="alert"]');
    expect(alert.text()).toContain("Capacity must be at least 1");
  });

  it("should emit the save event with valid inputs", async () => {
    const wrapper = mount(OrganiserEventForm, {
      props: {
        isEdit: false,
        event: null,
      },
    });

    await wrapper.find("#title").setValue("Valid Tech Workshop");
    await wrapper.find("#description").setValue("Learn Vue 3 with advanced Pinia.");

    const tomorrow = new Date();
    tomorrow.setDate(tomorrow.getDate() + 1);
    const formattedDate = tomorrow.toISOString().slice(0, 16);
    await wrapper.find("#dateTime").setValue(formattedDate);
    await wrapper.find("#capacity").setValue(75);
    await wrapper.find("#location").setValue("San Francisco");

    // Submit form
    await wrapper.find("form").trigger("submit.prevent");

    // Assert that save is emitted with full schema payload
    const emitted = wrapper.emitted("save");
    expect(emitted).toBeDefined();
    expect(emitted![0][0]).toMatchObject({
      title: "Valid Tech Workshop",
      description: "Learn Vue 3 with advanced Pinia.",
      location: "San Francisco",
      capacity: 75,
    });
  });
});
