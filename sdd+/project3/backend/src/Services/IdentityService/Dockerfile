FROM mcr.microsoft.com/dotnet/sdk:10.0 AS build
WORKDIR /src

# Copy csproj and restore dependencies
COPY ["IdentityService.csproj", "./"]
RUN dotnet restore "IdentityService.csproj"

# Copy everything else and build
COPY . .
RUN dotnet publish "IdentityService.csproj" -c Release -o /app/publish

FROM mcr.microsoft.com/dotnet/aspnet:10.0
WORKDIR /app
COPY --from=build /app/publish .
ENTRYPOINT ["dotnet", "IdentityService.dll"]
EXPOSE 80
