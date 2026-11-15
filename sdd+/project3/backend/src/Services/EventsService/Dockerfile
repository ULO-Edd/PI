FROM mcr.microsoft.com/dotnet/sdk:10.0 AS build
WORKDIR /src

# Copy csproj and restore dependencies
COPY ["EventsService.csproj", "./"]
RUN dotnet restore "EventsService.csproj"

# Copy everything else and build
COPY . .
RUN dotnet publish "EventsService.csproj" -c Release -o /app/publish

FROM mcr.microsoft.com/dotnet/aspnet:10.0
WORKDIR /app
COPY --from=build /app/publish .
ENTRYPOINT ["dotnet", "EventsService.dll"]
EXPOSE 80
