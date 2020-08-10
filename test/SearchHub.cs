// Generated with prosr1 by compiler 0.9, template 0.1

using Microsoft.AspNetCore.SignalR;
using Microsoft.AspNetCore.SignalR.Client;
using System;
using System.Threading.Tasks;

namespace Test.TestNamespace
{
	public interface ISearchHubClient : ISearchHubClientBase, IDisposable
	{
		Task InitializeHubConnection();
		Task CallSearchOnHub(SearchRequest message);
	}

	public interface ISearchHubClientBase
	{
		Task RecieveSearchRequestToAllAsync(SearchRequest message);
		Task RecieveSearchResponseToCallerAsync(SearchResponse message);
	}

		
	public abstract class SearchHubClientBase : ISearchHubClient
	{
		private readonly IHubConnectionBuilder _hubConnectionBuilder;

		private bool _disposed = false;
		private HubConnection _connection;

		protected SearchHubClientBase(IHubConnectionBuilder hubConnectionBuilder)
		{
			_hubConnectionBuilder = hubConnectionBuilder;
		}

		protected abstract Uri HubUrl { get; }

		public async Task InitializeHubConnection()
		{
			await GetConnection().ConfigureAwait(false);
		}
		
		public abstract Task RecieveSearchRequestToAllAsync(SearchRequest message);
		
		public abstract Task RecieveSearchResponseToCallerAsync(SearchResponse message);
		
		public async Task CallSearchOnHub(SearchRequest message)
		{
			var connection = await GetConnection().ConfigureAwait(false);

			await connection.SendAsync("Search", message);
		}

		private async Task<HubConnection> GetConnection()
		{
			if (_connection == null)
			{
				_connection = await CreateHubConnectionAsync().ConfigureAwait(false);
			}
		
			return _connection;
		}

		private async Task<HubConnection> CreateHubConnectionAsync()
		{
			var connection = _hubConnectionBuilder
				.WithUrl(HubUrl)
				.WithAutomaticReconnect(new[] { TimeSpan.Zero, TimeSpan.Zero, TimeSpan.FromSeconds(10) })
				.Build();

			BindClientMethods(ref connection);

			await connection
				.StartAsync()
				.ConfigureAwait(false);

			return connection;
		}

		private HubConnection BindClientMethods(ref HubConnection connection)
		{
			connection.On<SearchRequest>("RecieveSearchRequestToAllAsync", message => RecieveSearchRequestToAllAsync(message));
			connection.On<SearchResponse>("RecieveSearchResponseToCallerAsync", message => RecieveSearchResponseToCallerAsync(message));

			return connection;
		}

		public void Dispose()
		{
			Dispose(true);
			GC.SuppressFinalize(this);
		}
		
		protected virtual void Dispose(bool disposing)
		{
			if (!_disposed)
			{
				if (disposing)
				{
					if (!(_connection is null))
					{
						var result = _connection
							.DisposeAsync()
							.Wait(1000);

						if (!result)
						{
							throw new TimeoutException($"Timeout of one second reached while closing hub connection for {HubUrl}!");
						}
					}
				}

				_disposed = true;
			}
		}
		
		~SearchHubClientBase()
		{
			Dispose(false);
		}
	}

		
	public abstract class SearchHubBase : Hub<ISearchHubClientBase>
	{
		protected Task SendSearchRequestToAllAsync(SearchRequest message)
		{
			return Clients.All.RecieveSearchRequestToAllAsync(message);
		}

		public async Task Search(SearchRequest message)
		{
			var result = await HandleSearchAsync(message).ConfigureAwait(false);

			await Clients
				.Caller
				.RecieveSearchResponseToCallerAsync(result)
				.ConfigureAwait(false);
		}
		
		protected abstract Task<SearchResponse> HandleSearchAsync(SearchRequest message);
	}

	public class SearchRequest
	{
		public int result_per_page { get; set; }
		public int page_number { get; set; }
		public string query { get; set; }
	}

	public class SearchResponse
	{
		public Result result { get; set; }
	}

	public class Result
	{
		public string snippets { get; set; }
		public string title { get; set; }
		public string url { get; set; }
	}
}